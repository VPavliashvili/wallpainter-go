package args

func getHelp() boolarg {
	return boolarg{
		value:       &boolval{value: setValue(true)},
		names:       &[]string{"-h", "--help"},
		description: "print this help message",
	}
}

func getPath() stringarg {
	return stringarg{
		value:       &stringval{value: setValue("")},
		names:       &[]string{"-p", "--path"},
		description: "path of the directory where wallpapers are located in",
	}
}

