import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { UIShellModule } from 'carbon-components-angular';
import { AppComponent } from './app.component';
import { EventsModule } from './events/events.module';
import { HeaderComponent } from './header/header.component';
import { HomeModule } from './home/home.module';


describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RouterTestingModule, UIShellModule, EventsModule, HomeModule],
      declarations: [AppComponent, HeaderComponent],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });
});
