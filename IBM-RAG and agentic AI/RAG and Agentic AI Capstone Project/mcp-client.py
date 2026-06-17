
# Libraries for MCP client, LLM handling, and async operations
from dotenv import load_dotenv
import asyncio
import json
from pathlib import Path
from langchain_groq import ChatGroq
from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client
from mcp.types import (
Root,
TextContent, 
CreateMessageResult,
CreateMessageRequestParams,
)
import os
import getpass
# Load .env file explicitly with full path
load_dotenv(Path("/Users/anirudh/Desktop/Courseera/IBM-RAG and agentic AI/RAG and Agentic AI Capstone Project/.env"))
# Configuration
SERVER_SCRIPT = str(Path(__file__).parent/ "mcp-server.py")
PROJECT_DIR = Path(__file__).parent.resolve()

# StdioServerParameters launches "python server.py" via stdin/stdout
server_params = StdioServerParameters(
    command = "python3",
    args = [SERVER_SCRIPT] 
)

# ROOTS — Tell the server which directories it can access
def list_roots() -> list[Root]: 
    """Limit the server's file access to this project directory. """
    return [Root(uri = f"file://{PROJECT_DIR}",name = PROJECT_DIR.name)]

# ChatGroq client used to fulfill sampling requests from the server
ChatGroq_client = ChatGroq(
    model = "llama-3.3-70b-versatile", 
    max_tokens = 1024,
    temperature = 0.5
)

# SAMPLING — Handle LLM requests delegated from the server
async def handle_sampling(params: CreateMessageRequestParams) -> CreateMessageResult:
    """Run a Groq LLM call on behalf of the server and return the result."""
    # Extract the prompt text from the first sampling message
    prompt = params.messages[0].content.text

    print(f"\n[Sampling] Server requested LLM task:")
    print(f"  Prompt preview: {prompt[:150]}...")

    message = [{"role": "user", "content":prompt}] 
    response = ChatGroq_client.invoke(message)
    response_text = response.content
    return CreateMessageResult(
        role="assistant",
        content=TextContent(type="text", text=response_text),
        model="llama-3.3-70b-versatile",
    )

# HELPER — Open a session, call a tool, and return the parsed JSON result
async def call_tool(tool_name: str, arguments: dict) -> dict:
    """Connect to the server, call a tool, and return the parsed JSON result."""
    async with stdio_client(server_params) as (read, write):
        async with ClientSession(
            read,
            write,
            sampling_callback=handle_sampling,
            list_roots_callback=list_roots,
        ) as session:
            await session.initialize()
            result = await session.call_tool(tool_name, arguments=arguments)
            # The server always returns a single TextContent item
            return json.loads(result.content[0].text)

# CONNECTION & DISCOVERY
async def verify_connection():
    """Connect to the server and verify all expected tools and resources exist."""
    print("=" * 60)
    print("MCP Connection Verification")
    print("=" * 60)

    async with stdio_client(server_params) as (read, write):
        async with ClientSession(
            read,
            write,
            sampling_callback=handle_sampling,
            list_roots_callback=list_roots,
        ) as session:
            await session.initialize()

            # list_tools() sends a "tools/list" JSON-RPC request to the server
            tools_result = await session.list_tools()
            tool_names = [tool.name for tool in tools_result.tools]
            print("--- START SCREENSHOT ---")
            print(f"\nDiscovered {len(tool_names)} tools:")
            for tool in tools_result.tools:
                print(f"  - {tool.name}: {tool.description[:80]}...")

            assert "get_restaurant_info" in tool_names, "FAIL: get_restaurant_info not found!"
            assert "recommend_by_vibe" in tool_names, "FAIL: recommend_by_vibe not found!"
            assert "get_review" in tool_names, "FAIL: get_review not found!"
            print("\nAll required tools verified!")

            # list_resources() discovers data endpoints the server exposes
            resources_result = await session.list_resources()
            print(f"\nDiscovered {len(resources_result.resources)} resources:")
            for resource in resources_result.resources:
                print(f"  - {resource.uri}: {resource.name}")

            roots = list_roots()
            print(f"\nConfigured {len(roots)} roots:")
            for root in roots:
                print(f"  - {root.name}: {root.uri}")

            print("--- END SCREENSHOT ---")

# DEMOS — Call each tool through the MCP protocol
async def demo_get_restaurant_info():
    """Demo: Look up a restaurant by name."""
    print("\n" + "-" * 60)
    print("Demo: get_restaurant_info('Iron & Embers')")
    print("-" * 60)

    data = await call_tool("get_restaurant_info", {"restaurant_name": "Iron & Embers"})
    print(json.dumps(data, indent=2))

async def demo_recommend_by_vibe():
    """Demo: Find restaurants by vibe keyword."""
    print("\n" + "-" * 60)
    print("Demo: recommend_by_vibe('moody')")
    print("-" * 60)

    data = await call_tool("recommend_by_vibe", {"vibe": "moody"})
    print(f"Vibe: {data['vibe_searched']}")
    print(f"Structured matches: {len(data['structured_matches'])}")
    for match in data["structured_matches"]:
        print(f"  - {match['name']} ({match['cuisine']}) - {match['rating']}/5")
    print(f"Raw text excerpts: {len(data['raw_text_excerpts'])}")

async def demo_get_review():
    """Demo: Retrieve a restaurant review."""
    print("\n" + "-" * 60)
    print("Demo: get_review('Iron & Embers')")
    print("-" * 60)

    data = await call_tool("get_review", {"restaurant_name": "Iron & Embers"})
    print(json.dumps(data, indent=2))

# Main Entry Point
async def main():
    """Run all demos sequentially."""
    await demo_get_restaurant_info()   # ✅ replace first ...
    await demo_recommend_by_vibe()     # ✅ replace second ...
    await demo_get_review()            # ✅ replace third ...
    await verify_connection()

if __name__ == "__main__":
    asyncio.run(main())


