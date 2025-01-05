package message

type Model struct {
	ModelLoaded      bool   `json:"modelLoaded"`
	ModelName        string `json:"modelName"`
	ModelID          string `json:"modelID"`
	VtsModelName     string `json:"vtsModelName"`
	VtsModelIconName string `json:"vtsModelIconName"`
}

type ModelPosition struct {
	PositionX float32 `json:"position_x"`
	PositionY float32 `json:"position_y"`
	Rotation  float32 `json:"rotation"`
	Size      float32 `json:"size"`
}

type HotKey struct {
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	File             string   `json:"file"`
	HotkeyID         string   `json:"hotkeyID"`
	KeyCombination   []string `json:"keyCombination"`
	OnScreenButtonID int      `json:"onScreenButtonID"`
}

type UsedHotKey struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Parameters struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Expression struct {
	Name                       string       `json:"name"`
	File                       string       `json:"file"`
	Active                     bool         `json:"active"`
	DeactivateWhenKeyIsLetGo   bool         `json:"deactivateWhenKeyIsLetGo"`
	AutoDeactivateAfterSeconds bool         `json:"autoDeactivateAfterSeconds"`
	SecondsRemaining           uint16       `json:"secondsRemaining"`
	UsedInHotkeys              []UsedHotKey `json:"used_in_hotkeys"`
	Parameters                 []Parameters `json:"parameters"`
}

type Data struct {
	Active                      bool          `json:"active,omitempty"`
	VTubeStudioVersion          string        `json:"vTubeStudioVersion,omitempty"`
	CurrentSessionAuthenticated bool          `json:"currentSessionAuthenticated,omitempty"`
	Port                        uint16        `json:"port,omitempty"`
	InstanceID                  string        `json:"instanceID,omitempty"`
	WindowTitle                 string        `json:"windowTitle,omitempty"`
	PluginName                  string        `json:"pluginName,omitempty"`
	PluginDeveloper             string        `json:"pluginDeveloper,omitempty"`
	PluginIcon                  string        `json:"pluginIcon,omitempty"`
	AuthenticationToken         string        `json:"authenticationToken,omitempty"`
	ErrorID                     uint16        `json:"errorID,omitempty"`
	Message                     string        `json:"message,omitempty"`
	Reason                      string        `json:"reason,omitempty"`
	Uptime                      uint16        `json:"uptime,omitempty"`
	Framerate                   int           `json:"framerate,omitempty"`
	AllowedPlugins              uint16        `json:"allowedPlugins,omitempty"`
	ConnectedPlugins            uint16        `json:"connectedPlugins,omitempty"`
	StartedWithSteam            bool          `json:"startedWithSteam,omitempty"`
	WindowWidth                 uint16        `json:"windowWidth,omitempty"`
	WindowHeight                uint16        `json:"windowHeight,omitempty"`
	WindowIsFullscreen          bool          `json:"windowIsFullscreen,omitempty"`
	Models                      string        `json:"models,omitempty"`
	Backgrounds                 string        `json:"backgrounds,omitempty"`
	Items                       string        `json:"items,omitempty"`
	Config                      string        `json:"config,omitempty"`
	Logs                        string        `json:"logs,omitempty"`
	Backup                      string        `json:"backup,omitempty"`
	ModelLoaded                 bool          `json:"modelLoaded,omitempty"`
	ModelName                   string        `json:"modelName,omitempty"`
	ModelID                     string        `json:"modelID,omitempty"`
	VtsModelName                string        `json:"vtsModelName,omitempty"`
	VtsModelIconName            string        `json:"vtsModelIconName,omitempty"`
	Live2DModelName             string        `json:"live2DModelName,omitempty"`
	ModelLoadTime               uint16        `json:"modelLoadTime,omitempty"`
	TimeSinceModelLoaded        uint32        `json:"timeSinceModelLoaded,omitempty"`
	NumberOfLive2DParameters    uint16        `json:"numberOfLive2DParameters,omitempty"`
	NumberOfLive2DArtmeshes     uint16        `json:"numberOfLive2DArtmeshes,omitempty"`
	HasPhysicsFile              bool          `json:"hasPhysicsFile,omitempty"`
	NumberOfTextures            uint16        `json:"numberOfTextures,omitempty"`
	TextureResolution           uint16        `json:"textureResolution,omitempty"`
	ModelPosition               ModelPosition `json:"modelPosition,omitempty"`
	NumberOfModels              uint16        `json:"numberOfModels,omitempty"`
	AvailableModels             []Model       `json:"availableModels,omitempty"`
	TimeInSeconds               float32       `json:"timeInSeconds,omitempty"`
	ValuesAreRelativeToModel    bool          `json:"valuesAreRelativeToModel,omitempty"`
	PositionX                   float32       `json:"position_x,omitempty"`
	PositionY                   float32       `json:"position_y,omitempty"`
	Rotation                    float32       `json:"rotation,omitempty"`
	Size                        float32       `json:"size,omitempty"`
	Live2DItemFileName          string        `json:"live2DItemFileName,omitempty"`
	AvailableHotkeys            []HotKey      `json:"available_hotkeys,omitempty"`
	HotkeyID                    string        `json:"hotkeyID,omitempty"`
	ItemInstanceID              string        `json:"itemInstanceID,omitempty"`
	Details                     bool          `json:"details,omitempty"`
	ExpressionFile              string        `json:"expressionFile,omitempty"`
	Expressions                 []Expression  `json:"expressions,omitempty"`
	FadeTime                    float32       `json:"fadeTime,omitempty"`
	// Only finish until expression
	// Have not started ArtMeshes with color selection
	// Need to reorg this struct later, to be respond to request
	// for now I'm only using expression anyways
}

type Message struct {
	ApiName     string `json:"apiName"`
	ApiVersion  string `json:"apiVersion"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	RequestID   string `json:"requestID"`
	MessageType string `json:"messageType"`
	Data        Data   `json:"data,omitempty"`
}
