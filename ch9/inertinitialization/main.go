package main

import (
	"image"
	"sync"
)



var (
	// mu sync.Mutex
	// mu sync.RWMutex
	loadIconsOnce sync.Once
	icons map[string]image.Image
)
var test image.Image = image.NewRGBA(image.Rect(0, 0, 128, 128))


func loadIcon(_ string) image.Image{
	return test
}

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png": loadIcon("spades.png"),
		"hearts.png": loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png": loadIcon("club.png"),
	}
}

// // lazy initialization
// func Icon(name string) image.Image {
// 	if icons == nil {
// 		loadIcons()
// 	}

// 	return icons[name]
// }

// func Icon(name string) image.Image {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	if icons == nil {
// 		loadIcons()
// 	}

// 	return icons[name]
// }

// func Icon(name string) image.Image {
// 	mu.RLock()
// 	if icons != nil {
// 		icon := icons[name]
// 		mu.RUnlock()
// 		return icon
// 	}
// 	mu.RUnlock()

// 	mu.Lock()
// 	if icons != nil {
// 		loadIcons()
// 	}
// 	icon := icons[name]
// 	mu.Unlock()

// 	return icon
// }

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}