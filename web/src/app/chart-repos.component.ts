import { Component, OnInit } from '@angular/core';

import { ChartRepo } from './chart-repo';

import { ChartRepoService } from './chart-repo.service';

@Component({
  selector: 'chart-repos',
  templateUrl: './chart-repos.component.html',
  styleUrls: [ './chart-repos.component.css' ]
})
export class ChartReposComponent implements OnInit {

  repos: ChartRepo[] = [];
  selectedRepo: ChartRepo;

  constructor(private chartRepoService: ChartRepoService) { }

  ngOnInit(): void {
    this.chartRepoService.getRepos()
      .then(repos => this.repos = repos);
  }

  onSelect(chartRepo: ChartRepo): void {
    this.selectedRepo = chartRepo;
  }

  add(name: string, url: string): void {
    name = name.trim();
    if (!name) { return; }
    this.chartRepoService.create(name, url)
      .then(repo => {
        this.repos.push(repo);
        this.selectedRepo = null;
      });
  }

}
