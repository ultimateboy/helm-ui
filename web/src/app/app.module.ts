import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule }   from '@angular/forms';
import { HttpModule }    from '@angular/http';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent }         from './app.component';
import { ChartReposComponent }   from './chart-repos.component';
import { ChartRepoDetailComponent } from './chart-repo-detail.component';

import { ReleasesComponent }      from './releases.component';

import { ReleaseService }          from './release.service';
import { ChartRepoService }          from './chart-repo.service';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule
  ],
  declarations: [
    AppComponent,
    ChartReposComponent,
    ReleasesComponent,
    ChartRepoDetailComponent,
  ],
  providers: [ 
    ReleaseService,
    ChartRepoService
  ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }