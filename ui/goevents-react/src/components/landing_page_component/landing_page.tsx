import * as React from 'react';

export interface LandingPageProps {}

export class LandingPage extends React.Component<LandingPageProps, {}> {
    render() {
        return <div className="container">
            <h1>Welcome to the events site</h1>
            <p>Use the navigation bar at the top to view and book upcomming events</p>
        </div>
    }
}