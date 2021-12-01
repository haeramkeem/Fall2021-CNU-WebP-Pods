package config

var RootDir string = "/Users/haeramkeem/Documents/git/pods"

var config = map[string]string{
    "staticdir":    RootDir + "/ui/dist/",
}

func Get(query string) string {
    return config[query]
}
