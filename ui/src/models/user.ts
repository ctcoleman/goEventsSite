import { Booking } from './booking'

export interface User {
    ID: string
    First: string
    Last: string
    Age: string
    Bookings: Booking[]
}