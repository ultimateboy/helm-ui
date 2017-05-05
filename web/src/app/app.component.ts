import { Component } from '@angular/core';

@Component({
  selector: 'my-app',
  template: `
  <div class="title-bar">
    <a routerLink="/dashboard">
      <md-toolbar color="primary">
        <md-icon class="logo" color="accent">fingerprint</md-icon>
        <span class="title">{{title}}</span>
      </md-toolbar>
    </a>
  </div>
  <router-outlet></router-outlet>
  `,
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'HelmUI';
}
