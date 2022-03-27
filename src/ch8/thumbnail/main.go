package main

import (
	"log"
	"os"
	"sync"
)



func makeThumbnail1(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnail2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f) // Note: ignore errors
	}
}

func makeThubmnail3(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func (f string)  {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<- ch
	}
}

func makeThumbnail4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func (f string)  {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <- errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnail5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func (f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <- ch
		if it.err != nil {
			return nil, err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnail6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)

		go func (f string) {
			defer wg.Done()

			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}

			info, _ := os.Stat(thumb)

			sizes <- info.Size()
		}(f)
	}

	go func () {
		wg.Wait()
		close(sizes)
	}()

	var total int64

	for size := range sizes {
		total += size
	}

	return total
}