import os

from transformers import AutoTokenizer, AutoModelForSequenceClassification, pipeline
import socket


# Configs that use env
port = os.getenv("EMOTION_PORT")

# Model pipeline
classifier = pipeline("text-classification","/home/tomato/emotion-english-distilroberta-base", device=0)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM) # SOCK_STREAM -> TCP, SOCK_DGRAM -> UDP
address = ("localhost", port)
s.bind(address)

print("Starting socket")
while True:
    print("Waiting for connection")
    message = s.recv()