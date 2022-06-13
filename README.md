<p align="center">
	<img src="https://i.ibb.co/nQhm2pK/engine-DFM3-01.jpg" alt="engine-pict"/>
</p>


# Engine
Engine is just an http router, that implements middlewares.

This is not framework!! just http router
based on httprouter libs.

https://github.com/julienschmidt/httprouter

## Installation

use go modules to install.

```bash
go get github.com/vizucode/engine
```

## Usage

```go
//call NewEngine to initiate engine object.
router := engine.NewEngine()

router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Anya Forger.")
})

// to use a middleware, call engine.Use method
router.Use(SimpleMiddleware)

server := http.Server{
	Addr:    ":9000",
	Handler: router,
}

if err := server.ListenAndServe(); err != nil {
	panic(err)
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)