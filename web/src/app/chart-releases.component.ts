import { Component, OnInit, Input } from '@angular/core';
import { Router }            from '@angular/router';

import { Observable }        from 'rxjs/Observable';
import { Subject }           from 'rxjs/Subject';

// Observable class extensions
import 'rxjs/add/observable/of';

import { ChartRepoService } from './chart-repo.service';
import { ReleaseService } from './release.service';

import { Release } from './release';
import { Chart } from './chart';

@Component({
  selector: 'chart-releases',
  templateUrl: './chart-releases.component.html',
  styleUrls: [ './chart-releases.component.css' ],
  providers: [ReleaseService]
})
export class ChartReleasesComponent implements OnInit {
  @Input() chartName: string;
  @Input() repoName: string;

  releases: Release[];
  loading: boolean;

  constructor(
    private releaseService: ReleaseService,
    private chartRepoService: ChartRepoService,
    private router: Router) {}

  getChartReleases(chartName: string): void {
    this.releaseService.getChartReleases(chartName).then(releases => this.releases = releases);
  }

  ngOnInit(): void {
    this.getChartReleases(this.chartName);
  }

  toggleLoad(): void {
    this.loading = this.loading ? false : true;
  }

  install(): void {
    this.chartRepoService.install(this.chartName, this.repoName)
      .then(release => {
        if (this.releases) { 
          this.releases.push(release);
        } else {
            this.releases = [release];
        }
        this.router.navigate(['/dashboard']);
        
      });
  }

//   gotoDetail(hero: Hero): void {
//     let link = ['/detail', hero.id];
//     this.router.navigate(link);
//   }
}
