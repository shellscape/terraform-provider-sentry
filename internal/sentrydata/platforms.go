package sentrydata

// Platform represents a Sentry project platform type
type Platform string

// Platform constants based on https://github.com/getsentry/sentry/blob/master/src/sentry/models/project.py#L61-L170
const (
	PlatformOther                        Platform = "other"
	PlatformAndroid                      Platform = "android"
	PlatformApple                        Platform = "apple"
	PlatformAppleIOS                     Platform = "apple-ios"
	PlatformAppleMacOS                   Platform = "apple-macos"
	PlatformBun                          Platform = "bun"
	PlatformCapacitor                    Platform = "capacitor"
	PlatformCordova                      Platform = "cordova"
	PlatformDart                         Platform = "dart"
	PlatformDeno                         Platform = "deno"
	PlatformDotnet                       Platform = "dotnet"
	PlatformDotnetAspnet                 Platform = "dotnet-aspnet"
	PlatformDotnetAspnetcore             Platform = "dotnet-aspnetcore"
	PlatformDotnetAwslambda              Platform = "dotnet-awslambda"
	PlatformDotnetGcpfunctions           Platform = "dotnet-gcpfunctions"
	PlatformDotnetMaui                   Platform = "dotnet-maui"
	PlatformDotnetUwp                    Platform = "dotnet-uwp"
	PlatformDotnetWinforms               Platform = "dotnet-winforms"
	PlatformDotnetWpf                    Platform = "dotnet-wpf"
	PlatformDotnetXamarin                Platform = "dotnet-xamarin"
	PlatformElectron                     Platform = "electron"
	PlatformElixir                       Platform = "elixir"
	PlatformFlutter                      Platform = "flutter"
	PlatformGo                           Platform = "go"
	PlatformGoEcho                       Platform = "go-echo"
	PlatformGoFasthttp                   Platform = "go-fasthttp"
	PlatformGoFiber                      Platform = "go-fiber"
	PlatformGoGin                        Platform = "go-gin"
	PlatformGoHttp                       Platform = "go-http"
	PlatformGoIris                       Platform = "go-iris"
	PlatformGoMartini                    Platform = "go-martini"
	PlatformGoNegroni                    Platform = "go-negroni"
	PlatformGodot                        Platform = "godot"
	PlatformIonic                        Platform = "ionic"
	PlatformJava                         Platform = "java"
	PlatformJavaLog4j2                   Platform = "java-log4j2"
	PlatformJavaLogback                  Platform = "java-logback"
	PlatformJavaSpring                   Platform = "java-spring"
	PlatformJavaSpringBoot               Platform = "java-spring-boot"
	PlatformJavascript                   Platform = "javascript"
	PlatformJavascriptAngular            Platform = "javascript-angular"
	PlatformJavascriptAstro              Platform = "javascript-astro"
	PlatformJavascriptEmber              Platform = "javascript-ember"
	PlatformJavascriptGatsby             Platform = "javascript-gatsby"
	PlatformJavascriptNextjs             Platform = "javascript-nextjs"
	PlatformJavascriptNuxt               Platform = "javascript-nuxt"
	PlatformJavascriptReact              Platform = "javascript-react"
	PlatformJavascriptReactRouter        Platform = "javascript-react-router"
	PlatformJavascriptRemix              Platform = "javascript-remix"
	PlatformJavascriptSolid              Platform = "javascript-solid"
	PlatformJavascriptSolidstart         Platform = "javascript-solidstart"
	PlatformJavascriptSvelte             Platform = "javascript-svelte"
	PlatformJavascriptSveltekit          Platform = "javascript-sveltekit"
	PlatformJavascriptTanstackstartReact Platform = "javascript-tanstackstart-react"
	PlatformJavascriptVue                Platform = "javascript-vue"
	PlatformKotlin                       Platform = "kotlin"
	PlatformMinidump                     Platform = "minidump"
	PlatformNative                       Platform = "native"
	PlatformNativeQt                     Platform = "native-qt"
	PlatformNintendoSwitch               Platform = "nintendo-switch"
	PlatformNode                         Platform = "node"
	PlatformNodeAwslambda                Platform = "node-awslambda"
	PlatformNodeAzurefunctions           Platform = "node-azurefunctions"
	PlatformNodeCloudflarePages          Platform = "node-cloudflare-pages"
	PlatformNodeCloudflareWorkers        Platform = "node-cloudflare-workers"
	PlatformNodeConnect                  Platform = "node-connect"
	PlatformNodeExpress                  Platform = "node-express"
	PlatformNodeFastify                  Platform = "node-fastify"
	PlatformNodeGcpfunctions             Platform = "node-gcpfunctions"
	PlatformNodeHapi                     Platform = "node-hapi"
	PlatformNodeHono                     Platform = "node-hono"
	PlatformNodeKoa                      Platform = "node-koa"
	PlatformNodeNestjs                   Platform = "node-nestjs"
	PlatformPhp                          Platform = "php"
	PlatformPhpLaravel                   Platform = "php-laravel"
	PlatformPhpSymfony                   Platform = "php-symfony"
	PlatformPlaystation                  Platform = "playstation"
	PlatformPowershell                   Platform = "powershell"
	PlatformPython                       Platform = "python"
	PlatformPythonAiohttp                Platform = "python-aiohttp"
	PlatformPythonAsgi                   Platform = "python-asgi"
	PlatformPythonAwslambda              Platform = "python-awslambda"
	PlatformPythonBottle                 Platform = "python-bottle"
	PlatformPythonCelery                 Platform = "python-celery"
	PlatformPythonChalice                Platform = "python-chalice"
	PlatformPythonDjango                 Platform = "python-django"
	PlatformPythonFalcon                 Platform = "python-falcon"
	PlatformPythonFastapi                Platform = "python-fastapi"
	PlatformPythonFlask                  Platform = "python-flask"
	PlatformPythonGcpfunctions           Platform = "python-gcpfunctions"
	PlatformPythonPylons                 Platform = "python-pylons"
	PlatformPythonPymongo                Platform = "python-pymongo"
	PlatformPythonPyramid                Platform = "python-pyramid"
	PlatformPythonQuart                  Platform = "python-quart"
	PlatformPythonRq                     Platform = "python-rq"
	PlatformPythonSanic                  Platform = "python-sanic"
	PlatformPythonServerless             Platform = "python-serverless"
	PlatformPythonStarlette              Platform = "python-starlette"
	PlatformPythonTornado                Platform = "python-tornado"
	PlatformPythonTryton                 Platform = "python-tryton"
	PlatformPythonWsgi                   Platform = "python-wsgi"
	PlatformReactNative                  Platform = "react-native"
	PlatformRuby                         Platform = "ruby"
	PlatformRubyRack                     Platform = "ruby-rack"
	PlatformRubyRails                    Platform = "ruby-rails"
	PlatformRust                         Platform = "rust"
	PlatformUnity                        Platform = "unity"
	PlatformUnreal                       Platform = "unreal"
	PlatformXbox                         Platform = "xbox"
)

// AllPlatforms returns a slice of all valid Platform values
func AllPlatforms() []Platform {
	return []Platform{
		PlatformOther,
		PlatformAndroid,
		PlatformApple,
		PlatformAppleIOS,
		PlatformAppleMacOS,
		PlatformBun,
		PlatformCapacitor,
		PlatformCordova,
		PlatformDart,
		PlatformDeno,
		PlatformDotnet,
		PlatformDotnetAspnet,
		PlatformDotnetAspnetcore,
		PlatformDotnetAwslambda,
		PlatformDotnetGcpfunctions,
		PlatformDotnetMaui,
		PlatformDotnetUwp,
		PlatformDotnetWinforms,
		PlatformDotnetWpf,
		PlatformDotnetXamarin,
		PlatformElectron,
		PlatformElixir,
		PlatformFlutter,
		PlatformGo,
		PlatformGoEcho,
		PlatformGoFasthttp,
		PlatformGoFiber,
		PlatformGoGin,
		PlatformGoHttp,
		PlatformGoIris,
		PlatformGoMartini,
		PlatformGoNegroni,
		PlatformGodot,
		PlatformIonic,
		PlatformJava,
		PlatformJavaLog4j2,
		PlatformJavaLogback,
		PlatformJavaSpring,
		PlatformJavaSpringBoot,
		PlatformJavascript,
		PlatformJavascriptAngular,
		PlatformJavascriptAstro,
		PlatformJavascriptEmber,
		PlatformJavascriptGatsby,
		PlatformJavascriptNextjs,
		PlatformJavascriptNuxt,
		PlatformJavascriptReact,
		PlatformJavascriptReactRouter,
		PlatformJavascriptRemix,
		PlatformJavascriptSolid,
		PlatformJavascriptSolidstart,
		PlatformJavascriptSvelte,
		PlatformJavascriptSveltekit,
		PlatformJavascriptTanstackstartReact,
		PlatformJavascriptVue,
		PlatformKotlin,
		PlatformMinidump,
		PlatformNative,
		PlatformNativeQt,
		PlatformNintendoSwitch,
		PlatformNode,
		PlatformNodeAwslambda,
		PlatformNodeAzurefunctions,
		PlatformNodeCloudflarePages,
		PlatformNodeCloudflareWorkers,
		PlatformNodeConnect,
		PlatformNodeExpress,
		PlatformNodeFastify,
		PlatformNodeGcpfunctions,
		PlatformNodeHapi,
		PlatformNodeHono,
		PlatformNodeKoa,
		PlatformNodeNestjs,
		PlatformPhp,
		PlatformPhpLaravel,
		PlatformPhpSymfony,
		PlatformPlaystation,
		PlatformPowershell,
		PlatformPython,
		PlatformPythonAiohttp,
		PlatformPythonAsgi,
		PlatformPythonAwslambda,
		PlatformPythonBottle,
		PlatformPythonCelery,
		PlatformPythonChalice,
		PlatformPythonDjango,
		PlatformPythonFalcon,
		PlatformPythonFastapi,
		PlatformPythonFlask,
		PlatformPythonGcpfunctions,
		PlatformPythonPylons,
		PlatformPythonPymongo,
		PlatformPythonPyramid,
		PlatformPythonQuart,
		PlatformPythonRq,
		PlatformPythonSanic,
		PlatformPythonServerless,
		PlatformPythonStarlette,
		PlatformPythonTornado,
		PlatformPythonTryton,
		PlatformPythonWsgi,
		PlatformReactNative,
		PlatformRuby,
		PlatformRubyRack,
		PlatformRubyRails,
		PlatformRust,
		PlatformUnity,
		PlatformUnreal,
		PlatformXbox,
	}
}

// AllPlatformStrings returns a slice of all valid platform values as strings
func AllPlatformStrings() []string {
	platforms := AllPlatforms()
	result := make([]string, len(platforms))
	for i, p := range platforms {
		result[i] = string(p)
	}
	return result
}

// IsValid checks if the platform value is in the list of valid platforms
func (p Platform) IsValid() bool {
	switch p {
	case PlatformOther, PlatformAndroid, PlatformApple, PlatformAppleIOS, PlatformAppleMacOS,
		PlatformBun, PlatformCapacitor, PlatformCordova, PlatformDart, PlatformDeno,
		PlatformDotnet, PlatformDotnetAspnet, PlatformDotnetAspnetcore, PlatformDotnetAwslambda, PlatformDotnetGcpfunctions,
		PlatformDotnetMaui, PlatformDotnetUwp, PlatformDotnetWinforms, PlatformDotnetWpf, PlatformDotnetXamarin,
		PlatformElectron, PlatformElixir, PlatformFlutter, PlatformGo, PlatformGoEcho,
		PlatformGoFasthttp, PlatformGoFiber, PlatformGoGin, PlatformGoHttp, PlatformGoIris,
		PlatformGoMartini, PlatformGoNegroni, PlatformGodot, PlatformIonic, PlatformJava,
		PlatformJavaLog4j2, PlatformJavaLogback, PlatformJavaSpring, PlatformJavaSpringBoot, PlatformJavascript,
		PlatformJavascriptAngular, PlatformJavascriptAstro, PlatformJavascriptEmber, PlatformJavascriptGatsby, PlatformJavascriptNextjs,
		PlatformJavascriptNuxt, PlatformJavascriptReact, PlatformJavascriptReactRouter, PlatformJavascriptRemix, PlatformJavascriptSolid,
		PlatformJavascriptSolidstart, PlatformJavascriptSvelte, PlatformJavascriptSveltekit, PlatformJavascriptTanstackstartReact, PlatformJavascriptVue,
		PlatformKotlin, PlatformMinidump, PlatformNative, PlatformNativeQt, PlatformNintendoSwitch,
		PlatformNode, PlatformNodeAwslambda, PlatformNodeAzurefunctions, PlatformNodeCloudflarePages, PlatformNodeCloudflareWorkers,
		PlatformNodeConnect, PlatformNodeExpress, PlatformNodeFastify, PlatformNodeGcpfunctions, PlatformNodeHapi,
		PlatformNodeHono, PlatformNodeKoa, PlatformNodeNestjs, PlatformPhp, PlatformPhpLaravel,
		PlatformPhpSymfony, PlatformPlaystation, PlatformPowershell, PlatformPython, PlatformPythonAiohttp,
		PlatformPythonAsgi, PlatformPythonAwslambda, PlatformPythonBottle, PlatformPythonCelery, PlatformPythonChalice,
		PlatformPythonDjango, PlatformPythonFalcon, PlatformPythonFastapi, PlatformPythonFlask, PlatformPythonGcpfunctions,
		PlatformPythonPylons, PlatformPythonPymongo, PlatformPythonPyramid, PlatformPythonQuart, PlatformPythonRq,
		PlatformPythonSanic, PlatformPythonServerless, PlatformPythonStarlette, PlatformPythonTornado, PlatformPythonTryton,
		PlatformPythonWsgi, PlatformReactNative, PlatformRuby, PlatformRubyRack, PlatformRubyRails,
		PlatformRust, PlatformUnity, PlatformUnreal, PlatformXbox:
		return true
	}
	return false
}