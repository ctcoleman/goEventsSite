import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { HashRouter as Router, Route } from 'react-router-dom';
import { EventBookingFormContainer } from './components/event_booking_component/event_booking_form_container';
import { EventListContainer } from './components/event_list_component/event_list_container';
import { LandingPage } from './components/landing_page_component/landing_page';
import { Navigation } from './components/navigation_component/navigation';

class App extends React.Component<{}, {}> {
    render() {
        const eventList = () => <EventListContainer eventListURL="http://localhost:8888/events" />;
        const eventBooking = ({match}: any) =>
            <EventBookingFormContainer
                eventID={match.params.id}
                eventServiceURL="http://localhost:8888/"
                bookingServiceURL="http://localhost:6666" />;
        
        return <Router>
            <div className="container">
                <Navigation brandName="GoEvents" />
                <h1>GoEvents Site</h1>
                <Route exact path="/" component={LandingPage}/>
                <Route path="/events" component={eventList} />
                <Route path="/events/:id/book" component={eventBooking} />
            </div>
        </Router>
    }
}

ReactDOM.render(
    <App />,
    document.getElementById("goevents-app")
)