import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ChartReposComponent }   from './chart-repos.component';
import { ReleasesComponent }      from './releases.component';
import { ChartRepoDetailComponent } from './chart-repo-detail.component';

const routes: Routes = [
  { path: 'dashboard', redirectTo: '', pathMatch: 'full' },
  { path: 'chart-repos',  component: ChartReposComponent },
  { path: 'chart-repos/detail/:name', component: ChartRepoDetailComponent },
  { path: '',     component: ReleasesComponent }
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
