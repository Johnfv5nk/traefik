package file

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"sync"
)

type Provider struct {
	watcher *fsnotify.Watcher
	mu      sync.Mutex
	watched map[string]bool
}

func (p *Provider) WatchCertificate(path string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.watched[path] {
		return
	}

	err := p.watcher.Add(path)
	if err != nil {
		log.Printf("failed to watch certificate file %s: %v", path, err)
		return
	}
	p.watched[path] = true
}

func (p *Provider) HandleEvent(event fsnotify.Event) {
	// Trigger reload logic when certificate file changes
	if event.Op&fsnotify.Write == fsnotify.Write {
		log.Printf("certificate file modified: %s, triggering reload", event.Name)
		// Trigger configuration reload here
	}
}