import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule }   from '@angular/forms';
import { HttpModule }    from '@angular/http';

//import {MdButtonModule} from '@angular/material';
import {MaterialModule} from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent }         from './app.component';
import { ChartReposComponent }   from './chart-repos.component';
import { ChartRepoDetailComponent } from './chart-repo-detail.component';

import { ReleasesComponent }      from './releases.component';
import { ChartReleasesComponent } from './chart-releases.component'

import { ReleaseService }          from './release.service';
import { ChartRepoService }          from './chart-repo.service';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule,
    MaterialModule.forRoot(),
    BrowserAnimationsModule
  ],
  declarations: [
    AppComponent,
    ChartReposComponent,
    ReleasesComponent,
    ChartRepoDetailComponent,
    ChartReleasesComponent,
  ],
  providers: [ 
    ReleaseService,
    ChartRepoService
  ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }