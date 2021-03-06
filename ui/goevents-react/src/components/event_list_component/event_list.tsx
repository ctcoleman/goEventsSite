import * as React from 'react';
import { Event } from '../../models/event';
import { EventListItem } from './event_list_item';

export interface EventListProps {
    events: Event[];
}

export class EventList extends React.Component<EventListProps, {}> {
    render() {
        const items = this.props.events.map(e =>
            <EventListItem event={e} key={e.ID}/>
        );
        return <table className="bx--data-table">
            <thead>
                <tr>
                    <th>Event</th>
                    <th>Where</th>
                    <th>Start</th>
                    <th>End</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {items}
            </tbody>
        </table>
    }
}