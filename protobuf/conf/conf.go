package conf

func (a *App) IsDevelopment() bool {
	return a.Env == "dev"
}
