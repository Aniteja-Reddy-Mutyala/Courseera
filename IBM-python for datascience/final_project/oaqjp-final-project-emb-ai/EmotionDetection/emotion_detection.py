import requests
import json

def emotion_detector(text_to_analyse):
    """
    Analyzes the emotion in the provided text using Watson NLP library.
    Handles blank entries and API errors.
    
    Args:
        text_to_analyse (str): The text to be analyzed for emotions
        
    Returns:
        dict: Dictionary containing emotion scores and dominant emotion,
              or None values if input is blank or request fails
    """
    # URL for Watson NLP Emotion Predict function
    url = 'https://sn-watson-emotion.labs.skills.network/v1/watson.runtime.nlp.v1/NlpService/EmotionPredict'
    
    # Headers required for the API request
    headers = {"grpc-metadata-mm-model-id": "emotion_aggregated-workflow_lang_en_stock"}
    
    # Input JSON format with the text to analyze
    input_json = {"raw_document": {"text": text_to_analyse}}
    
    # Make POST request to Watson NLP library
    response = requests.post(url, json=input_json, headers=headers)
    
    # Check status code for blank entries
    if response.status_code == 400:
        # Return dictionary with None values for all keys
        return {
            'anger': None,
            'disgust': None,
            'fear': None,
            'joy': None,
            'sadness': None,
            'dominant_emotion': None
        }
    
    # Convert response text to dictionary
    formatted_response = json.loads(response.text)
    
    # Extract emotion scores from the response
    emotions = formatted_response['emotionPredictions'][0]['emotion']
    
    # Extract individual emotion scores
    anger_score = emotions['anger']
    disgust_score = emotions['disgust']
    fear_score = emotions['fear']
    joy_score = emotions['joy']
    sadness_score = emotions['sadness']
    
    # Find the dominant emotion (emotion with highest score)
    emotion_scores = {
        'anger': anger_score,
        'disgust': disgust_score,
        'fear': fear_score,
        'joy': joy_score,
        'sadness': sadness_score
    }
    
    # Get the dominant emotion
    dominant_emotion = max(emotion_scores, key=emotion_scores.get)
    
    # Return formatted output
    return {
        'anger': anger_score,
        'disgust': disgust_score,
        'fear': fear_score,
        'joy': joy_score,
        'sadness': sadness_score,
        'dominant_emotion': dominant_emotion
    }