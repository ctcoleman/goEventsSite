// Event interface describes the format for an Event item
export interface Event {
  ID: string;
  Name: string;
  StartDate: number;
  EndDate: number;
  Duration: number;
  Location: {
    ID: string;
    Name: string;
    City: string;
    State: string;
    Country: string;
    OpenTime: number;
    CloseTime: number;
  };
}
