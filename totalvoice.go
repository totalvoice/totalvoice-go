package totalvoice

// BaseURI - URI Base
const (
	BaseURI = "https://api2.totalvoice.com.br"
)

// HTTPClient Interface for Clients
type HTTPClient interface {
	GetResource(path string) (string, error)
	CreateResource(values map[string]string, path string) (string, error)
	UpdateResource(values map[string]string, path string) (string, error)
	DeleteResource(path string) (string, error)
	SetBaseURI(value string)
	GetBaseURI() string
}

// TotalVoice struct
type TotalVoice struct {
	client HTTPClient

	Perfil *Perfil
	Audio  *Audio
}

// NewTotalVoiceClient - Cria TotalVoice struct.
func NewTotalVoiceClient(accessToken string) *TotalVoice {

	c := &Client{accessToken: accessToken, baseURI: BaseURI}

	tvce := &TotalVoice{client: c}

	tvce.Perfil = &Perfil{client: c}
	tvce.Audio = &Audio{client: c}

	return tvce
}

// NewClient - Cria TotalVoice struct.
func NewClient(accessToken string, baseURI string) *TotalVoice {

	tvce := NewTotalVoiceClient(accessToken)
	tvce.client.SetBaseURI(baseURI)

	return tvce
}
