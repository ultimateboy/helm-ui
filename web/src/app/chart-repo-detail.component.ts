// Keep the Input import for now, you'll remove it later:
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params }   from '@angular/router';
import { Location }                 from '@angular/common';

import { ChartRepoService } from './chart-repo.service';
import { ChartRepo } from './chart-repo'
import { Chart } from './chart'

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'chart-repo-detail',
  templateUrl: './chart-repo-detail.component.html',
  styleUrls: [ './chart-repo-detail.component.css' ]
})

export class ChartRepoDetailComponent implements OnInit {

    repo: string;
    charts: Chart[];

    constructor(
      private chartRepoService: ChartRepoService,
      private route: ActivatedRoute,
      private location: Location
    ) {}

    ngOnInit(): void {
      this.route.params
        .switchMap((params: Params) => this.chartRepoService.getRepoCharts(params['name']))
        .subscribe(charts => this.charts = charts);
      this.repo = this.route.snapshot.params['name'];
    }

    install(chart: Chart): void {
      this.chartRepoService.install(chart.name, this.repo);
    }

    goBack(): void {
      this.location.back();
    }
}