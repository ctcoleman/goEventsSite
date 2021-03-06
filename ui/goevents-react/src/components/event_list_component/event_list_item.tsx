import * as React from 'react';
import { Link } from 'react-router-dom';
import { Event } from '../../models/event';

export interface EventListItemProps {
    event: Event;
}

export class EventListItem extends React.Component<EventListItemProps, {}> {
    render() {
        const start = new Date(this.props.event.StartDate * 1000);
        const end = new Date(this.props.event.EndDate * 1000);

        return <tr>
            <td>{this.props.event.Name}</td>
            <td>{this.props.event.Location.Name}</td>
            <td>{start.toLocaleDateString()}</td>
            <td>{end.toLocaleDateString()}</td>
            <td>
                <Link to={`/events/${this.props.event.ID}/book`}>Book that shit!</Link>
            </td>
        </tr>
    }
}