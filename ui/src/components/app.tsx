import * as React from 'react'
import { EventListContainer } from './event_list_component/event_list_container'


export class App extends React.Component<{}, {}> {
    render() {
        return <div className="container">
            <header>
                <h1>GoEvents</h1>
            </header>
            <div className="eventsList">
                <EventListContainer eventListURL="http://localhost:8181"/>
            </div>
        </div>
    }
}