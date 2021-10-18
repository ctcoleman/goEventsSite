import {Location} from './location'

export interface Event {
    ID: string
    Name: string
    StartDate: number
    EndDate: number
    Duration: number
    Location: Location
}