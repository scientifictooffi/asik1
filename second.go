package main

import "fmt"

type Observer interface {
	Update(serverName string, status string)
}
type ServerStatusObserver struct {
	name string
}

func (s *ServerStatusObserver) Update(serverName string, status string) {
	fmt.Printf("Наблюдатель %s получил обновление: Сервер %s - Статус: %s\n", s.name, serverName, status)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(serverName string, status string)
}
type ServerMonitor struct {
	observers []Observer
}

func (s *ServerMonitor) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ServerMonitor) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ServerMonitor) NotifyObservers(serverName string, status string) {
	for _, observer := range s.observers {
		observer.Update(serverName, status)
	}
}

func (s *ServerMonitor) ChangeServerStatus(serverName string, status string) {
	fmt.Printf("Изменение статуса сервера %s: %s\n", serverName, status)
	s.NotifyObservers(serverName, status)
}

func main() {
	serverMonitor := &ServerMonitor{}

	observer1 := &ServerStatusObserver{name: "Observer 1"}
	observer2 := &ServerStatusObserver{name: "Observer 2"}
	observer3 := &ServerStatusObserver{name: "Observer 3"}
	serverMonitor.RegisterObserver(observer1)
	serverMonitor.RegisterObserver(observer2)
	serverMonitor.RegisterObserver(observer3)

	serverMonitor.ChangeServerStatus("Server 1", "Online")

	serverMonitor.RemoveObserver(observer2)

	serverMonitor.ChangeServerStatus("Server 2", "Offline")
}
