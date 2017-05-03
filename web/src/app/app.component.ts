import { Component } from '@angular/core';

@Component({
  selector: 'my-app',
  template: `
  <md-sidenav-container [class.m2app-dark]="isDarkTheme">

  <md-sidenav #sidenav mode="over" class="app-sidenav" color="primary">
    <md-toolbar>
      <span>Options</span>
      <md-toolbar-row>
        <a routerLink="/chart-repos" (click)="sidenav.toggle()" routerLinkActive="active"><button class="app-button" md-button>Chart Repos</button></a>
      </md-toolbar-row>
      <md-toolbar-row>
        <a routerLink="/releases" (click)="sidenav.toggle()" routerLinkActive="active"><button class="app-button" md-button>Releases</button></a>
      </md-toolbar-row>
    </md-toolbar>
   </md-sidenav>

  <md-toolbar color="primary">
    <button md-button class="app-menu-button" (click)="sidenav.toggle()">
      <i class="material-icons">menu</i>
    </button>
    <span class="title">{{title}}</span>
  </md-toolbar>
    <router-outlet></router-outlet>
    </md-sidenav-container>
  `,
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'HelmUI';
}
