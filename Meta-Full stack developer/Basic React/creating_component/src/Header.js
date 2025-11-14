const handleClick=()=><p style={{borderColor:"blue"}}>You have clicked me</p>
function Header({title,color}){
  return (
  <>
  <h1 style={{color:color}}>{title}</h1>
  <button onClick={handleClick}> Click me please</button>
  </>
)}

export default Header