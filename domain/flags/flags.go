package flags

const (
	Help         = "--help"
	SetWallpaper = "--set-wallpaper"
)

type Flag string

func ToFlag(s string) Flag {
    return Flag(s)
}
