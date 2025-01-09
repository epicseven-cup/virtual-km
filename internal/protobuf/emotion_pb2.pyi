from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Emotion(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    sadness: _ClassVar[Emotion]
    neutral: _ClassVar[Emotion]
    surprise: _ClassVar[Emotion]
    disgust: _ClassVar[Emotion]
    joy: _ClassVar[Emotion]
    fear: _ClassVar[Emotion]
sadness: Emotion
neutral: Emotion
surprise: Emotion
disgust: Emotion
joy: Emotion
fear: Emotion

class EmotionMessage(_message.Message):
    __slots__ = ("text",)
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class EvaluatedEmotionMessage(_message.Message):
    __slots__ = ("text", "emotion")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    EMOTION_FIELD_NUMBER: _ClassVar[int]
    text: str
    emotion: Emotion
    def __init__(self, text: _Optional[str] = ..., emotion: _Optional[_Union[Emotion, str]] = ...) -> None: ...
