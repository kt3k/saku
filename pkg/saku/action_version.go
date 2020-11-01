package saku

func actionVersion() error {
	colorablePrintf("saku@%s\n", Version)

	return nil
}
