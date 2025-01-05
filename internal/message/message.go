package message

type Model struct {
	modelLoaded      bool   `json:"modelLoaded"`
	modelName        string `json:"modelName"`
	modelID          string `json:"modelID"`
	vtsModelName     string `json:"vtsModelName"`
	vtsModelIconName string `json:"vtsModelIconName"`
}

type ModelPosition struct {
	positionX float32 `json:"position_x"`
	positionY float32 `json:"position_y"`
	rotation  float32 `json:"rotation"`
	size      float32 `json:"size"`
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
	name string `json:"name"`
	id   string `json:"id"`
}

type Parameters struct {
	name  string `json:"name"`
	value int    `json:"value"`
}

type Expression struct {
	Name                       string       `json:"name"`
	File                       string       `json:"file"`
	Active                     bool         `json:"active"`
	deactivateWhenKeyIsLetGo   bool         `json:"deactivateWhenKeyIsLetGo"`
	autoDeactivateAfterSeconds bool         `json:"autoDeactivateAfterSeconds"`
	secondsRemaining           uint16       `json:"secondsRemaining"`
	usedInHotkeys              []UsedHotKey `json:"used_in_hotkeys"`
	parameters                 []Parameters `json:"parameters"`
}

type Data struct {
	active                      bool          `json:"active,omitempty"`
	vTubeStudioVersion          string        `json:"vTubeStudioVersion,omitempty"`
	currentSessionAuthenticated bool          `json:"currentSessionAuthenticated,omitempty"`
	port                        uint16        `json:"port,omitempty"`
	instanceID                  string        `json:"instanceID,omitempty"`
	windowTitle                 string        `json:"windowTitle,omitempty"`
	pluginName                  string        `json:"pluginName,omitempty"`
	pluginDeveloper             string        `json:"pluginDeveloper,omitempty"`
	pluginIcon                  string        `json:"pluginIcon,omitempty"`
	authenticationToken         string        `json:"authenticationToken,omitempty"`
	errorID                     uint16        `json:"errorID,omitempty"`
	message                     string        `json:"message,omitempty"`
	reason                      string        `json:"reason,omitempty"`
	uptime                      uint16        `json:"uptime,omitempty"`
	framerate                   int           `json:"framerate,omitempty"`
	allowedPlugins              uint16        `json:"allowedPlugins,omitempty"`
	connectedPlugins            uint16        `json:"connectedPlugins,omitempty"`
	startedWithSteam            bool          `json:"startedWithSteam,omitempty"`
	windowWidth                 uint16        `json:"windowWidth,omitempty"`
	windowHeight                uint16        `json:"windowHeight,omitempty"`
	windowIsFullscreen          bool          `json:"windowIsFullscreen,omitempty"`
	models                      string        `json:"models,omitempty"`
	backgrounds                 string        `json:"backgrounds,omitempty"`
	items                       string        `json:"items,omitempty"`
	config                      string        `json:"config,omitempty"`
	logs                        string        `json:"logs,omitempty"`
	backup                      string        `json:"backup,omitempty"`
	modelLoaded                 bool          `json:"modelLoaded,omitempty"`
	modelName                   string        `json:"modelName,omitempty"`
	modelID                     string        `json:"modelID,omitempty"`
	vtsModelName                string        `json:"vtsModelName,omitempty"`
	vtsModelIconName            string        `json:"vtsModelIconName,omitempty"`
	live2DModelName             string        `json:"live2DModelName,omitempty"`
	modelLoadTime               uint16        `json:"modelLoadTime,omitempty"`
	timeSinceModelLoaded        uint32        `json:"timeSinceModelLoaded,omitempty"`
	numberOfLive2DParameters    uint16        `json:"numberOfLive2DParameters,omitempty"`
	numberOfLive2DArtmeshes     uint16        `json:"numberOfLive2DArtmeshes,omitempty"`
	hasPhysicsFile              bool          `json:"hasPhysicsFile,omitempty"`
	numberOfTextures            uint16        `json:"numberOfTextures,omitempty"`
	textureResolution           uint16        `json:"textureResolution,omitempty"`
	modelPosition               ModelPosition `json:"modelPosition,omitempty"`
	numberOfModels              uint16        `json:"numberOfModels,omitempty"`
	availableModels             []Model       `json:"availableModels,omitempty"`
	timeInSeconds               float32       `json:"timeInSeconds,omitempty"`
	valuesAreRelativeToModel    bool          `json:"valuesAreRelativeToModel,omitempty"`
	positionX                   float32       `json:"position_x,omitempty"`
	positionY                   float32       `json:"position_y,omitempty"`
	rotation                    float32       `json:"rotation,omitempty"`
	size                        float32       `json:"size,omitempty"`
	live2DItemFileName          string        `json:"live2DItemFileName,omitempty"`
	availableHotkeys            []HotKey      `json:"available_hotkeys,omitempty"`
	hotkeyID                    string        `json:"hotkeyID,omitempty"`
	itemInstanceID              string        `json:"itemInstanceID,omitempty"`
	details                     bool          `json:"details,omitempty"`
	expressionFile              string        `json:"expressionFile,omitempty"`
	expressions                 []Expression  `json:"expressions,omitempty"`
	fadeTime                    float32       `json:"fadeTime,omitempty"`
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
