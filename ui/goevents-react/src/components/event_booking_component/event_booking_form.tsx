import * as React from 'react';
import { ChangeEvent } from 'react';
import { Event } from '../../models/event';
import { FormRow } from './event_booking_form_row';

export interface EventBookingFormProps {
    event: Event;
    onSubmit: (seats: number) => any
}

export interface EventBookingFormState {
    seats: number;
}

export class EventBookingForm extends React.Component<EventBookingFormProps, EventBookingFormState> {
    constructor(p: EventBookingFormProps) {
        super(p);

        this.state = { seats: 1 };
    }

    private handleNewAmount(event: ChangeEvent<HTMLSelectElement>) {
        const newState: EventBookingFormState = {
            seats: parseInt(event.target.value)
        };

        this.setState(newState);
    }

    render() {
        return <div>
            <h2>Book your tickets for {this.props.event.Name}</h2>
            <form>
                <FormRow label="Event">
                    <p className="form-control-static">
                        {this.props.event.Name}
                    </p>
                </FormRow>
                <FormRow label="Number of tickets">
                    {/* <div className="bx--number__input-wrapper">
                        <input className="form-control" type="number" min="0" max="10" value={this.state.seats} onChange={event => this.handleNewAmount(event)}>
                            <div className="bx--number__controls">
                                <button className="bx--number__control-btn up-icon" type="button">
                                    <svg focusable="false" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg" width="8" height="4" viewBox="0 0 8 4" aria-hidden="true"><path d="M0 4L4 0 8 4z"></path></svg>
                                </button>
                                <button className="bx--number__control-btn down-icon" type="button">
                                    <svg focusable="false" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg" width="8" height="4" viewBox="0 0 8 4" aria-hidden="true"><path d="M8 0L4 4 0 0z"></path></svg>
                                </button>
                            </div>
                        </input>
                    </div> */}
                    <div className="bx--select">
                        <label htmlFor="select-id" className="bx--label">Tickets</label>
                        <div className="bx--select-input__wrapper">
                            <select id="select-id" className="bx--select-input" value={this.state.seats} onChange={e => this.handleNewAmount(e)}>
                                <option className="bx--select-option" value="1">1</option>
                                <option value="2">2</option>
                                <option value="3">3</option>
                                <option value="4">4</option>
                                <option value="5">5</option>
                            </select>
                        </div>
                    </div>
                </FormRow>
                <FormRow>
                    <button 
                        className="bx--btn bx--btn--primary"
                        type="button"
                        onClick={() => this.props.onSubmit(this.state.seats)}>Submit Order
                    </button>
                </FormRow>
            </form>
        </div>
    }
}