package asset

type Source struct {
	Name                      string
	Path                      string
	AdminEsbuildCompatible    bool
	FrontendEsbuildCompatible bool
	DisableSass               bool
	NpmStrict                 bool
}
