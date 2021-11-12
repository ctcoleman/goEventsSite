import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EventsListContainerComponent } from './events-list-container.component';

describe('EventsListContainerComponent', () => {
  let component: EventsListContainerComponent;
  let fixture: ComponentFixture<EventsListContainerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EventsListContainerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EventsListContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
