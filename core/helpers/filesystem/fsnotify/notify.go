package fsnotify

import (
	"github.com/fsnotify/fsnotify"
	"github.com/kyaxcorp/go-core/core/helpers/err"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type Op fsnotify.Op

// Just extending!
// These are the operations of an event!
const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

// This is the Event structure when we receive from the watcher
type EventData struct {
	// This is the path where it happened!
	Path string
	// This is the time when happened the event
	Time time.Time
	// This is the filtered Name
	Name string
	// This is the operation that happened
	Op Op
}

type WatchCallback func(e EventData)

// This is the structure when we send the configuration to the watcher!
type WatchConfig struct {
	// This is the callback which should be called!
	Callback WatchCallback
	// What operations should listen on!
	Op []Op
}

type OnError func(err error)

type Notifier struct {
	// The watcher
	watcher *fsnotify.Watcher

	// The paths that are being watched
	watchPaths map[string]WatchConfig

	status bool

	// If is there any path in the stack
	isAnyPathAdded bool

	// On Error Callback
	onError OnError
}

func New() (*Notifier, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Notifier{
		// Attach the watcher
		watcher: watcher,
		// Create the map!
		watchPaths: make(map[string]WatchConfig),
	}, nil
}

func (n *Notifier) IsStarted() bool {
	return n.status
}

func (n *Notifier) OnError(callback OnError) *Notifier {
	// Even nil is accepted!
	n.onError = callback
	return n
}

func (n *Notifier) IsStopped() bool {
	return !n.status
}

func (n *Notifier) checkNrOfWatchPaths() {
	if len(n.watchPaths) > 0 {
		n.isAnyPathAdded = true
	} else {
		n.isAnyPathAdded = false
	}
}

func (n *Notifier) Watch(
	path string,
	watchConfig WatchConfig,
) *Notifier {

	if path != "" {
		path, _ = filesystem.RealPath(path)
	}

	// Add the path
	if path == "" || watchConfig.Callback == nil {
		err.New(0, "FsNotify - Path empty or callback")
		return n
	}

	// Check if exists already!
	//We can overwrite the existing callback!

	if _, ok := n.watchPaths[path]; !ok {
		// Exists!
		// Add to watcher
		_err := n.watcher.Add(path)
		if _err != nil {
			err.New(0, _err.Error())
			return n
		}
	}
	// Write or overwrite the callback!
	n.watchPaths[path] = watchConfig
	n.checkNrOfWatchPaths()
	return n
}

// Remove the watcher!
func (n *Notifier) Remove(path string) *Notifier {
	if _, ok := n.watchPaths[path]; ok {
		_err := n.watcher.Remove(path)
		if _err != nil {
			// log.Println("Failed removing", _err)
			err.New(0, "Failed removing -> "+_err.Error())
			return n
		}
		delete(n.watchPaths, path)
		n.checkNrOfWatchPaths()
	}
	return n
}

func (n *Notifier) Start() *Notifier {
	if n.IsStarted() {
		log.Println("notifier already started")
		return n
	}

	go func() {
		// Entering infinite loop
		// TODO: we should do a finished variable!
		for {
			select {
			case event, ok := <-n.watcher.Events:
				// Getting events through Channel
				if !ok {
					return
				}

				// If are any paths...
				if !n.isAnyPathAdded {
					return
				}

				for path, watchConfig := range n.watchPaths {
					if !strings.Contains(event.Name, path) {
						continue
					}

					runCallback := false
					if len(watchConfig.Op) > 0 {
						for _, op := range watchConfig.Op {
							if Op(event.Op) == op {
								runCallback = true
								break
							}
						}
					} else {
						// It's everything! Listening to all!
						runCallback = true
					}

					if !runCallback {
						// Skipping this path which is being watched
						continue
					}

					//log.Println("aaa", event.Op, Op(event.Op))

					watchConfig.Callback(EventData{
						Path: event.Name,
						Time: time.Now(),
						Name: filepath.Base(event.Name), // Filter the file Name!
						Op:   Op(event.Op),
					})
				}

				/*log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}*/
			}
		}

		log.Println("Notifier goroutine stopped!")
	}()

	// On Errors
	go func() {
		// Entering infinite loop
		// TODO: we should do a finished variable!
		for {
			select {
			case err, ok := <-n.watcher.Errors:
				// Getting events through Channel
				if !ok {
					return
				}

				// Callback!
				if n.onError != nil {
					n.onError(err)
				}
				log.Println("error:", err)
			}
		}

		log.Println("Notifier goroutine stopped!")
	}()

	return n
}

func (n *Notifier) Stop() *Notifier {
	if n.IsStopped() {
		log.Println("notifier already stopped!")
		return n
	}

	// TODO: stop the created goroutines!
	n.watcher.Close()
	log.Println("notifier stopped")
	return n
}
