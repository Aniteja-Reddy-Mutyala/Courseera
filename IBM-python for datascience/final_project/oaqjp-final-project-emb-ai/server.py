"""
Flask server for Emotion Detection application.

This module provides a web interface and API endpoint for emotion analysis
using the Watson NLP library through the EmotionDetection package.
"""

from flask import Flask, render_template, request
from EmotionDetection import emotion_detector

# Initialize Flask application
app = Flask("Emotion Detector")


@app.route("/emotionDetector")
def emotion_analyzer():
    """
    Analyze emotion from text provided via GET parameter.

    Receives text through 'textToAnalyze' parameter and returns
    a formatted response with emotion scores and dominant emotion.
    Handles blank entries with appropriate error message.

    Returns:
        str: Formatted string with emotion analysis or error message
    """
    # Get the text to analyze from request arguments
    text_to_analyze = request.args.get('textToAnalyze')

    # Call the emotion_detector function
    response = emotion_detector(text_to_analyze)

    # Check if dominant_emotion is None (blank entry or error)
    if response['dominant_emotion'] is None:
        return "Invalid text! Please try again!"

    # Extract emotion scores
    anger = response['anger']
    disgust = response['disgust']
    fear = response['fear']
    joy = response['joy']
    sadness = response['sadness']
    dominant_emotion = response['dominant_emotion']

    # Format the response string
    formatted_response = (
        f"For the given statement, the system response is "
        f"'anger': {anger}, 'disgust': {disgust}, 'fear': {fear}, "
        f"'joy': {joy} and 'sadness': {sadness}. "
        f"The dominant emotion is {dominant_emotion}."
    )

    return formatted_response


@app.route("/")
def render_index_page():
    """
    Render the main application page.

    Returns:
        str: Rendered HTML template for the index page
    """
    return render_template('index.html')


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)
