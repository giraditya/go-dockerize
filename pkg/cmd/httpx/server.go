package httpx

func RunHttp() {
	app := setupRouter()
	app.Run(":3000")
}
