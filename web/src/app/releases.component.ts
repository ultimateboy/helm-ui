import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Release } from './release';
import { ReleaseService } from './release.service';

@Component({
  selector: 'my-releases',
  templateUrl: './releases.component.html',
  styleUrls: [ './releases.component.css' ]
})
export class ReleasesComponent implements OnInit {
  releases: Release[];
  selectedRelease: Release;

  constructor(
    private router: Router,
    private releaseService: ReleaseService) { }

  getReleases(): void {
    this.releaseService.getReleases().then(releases => this.releases = releases);
  }

  ngOnInit(): void {
    this.getReleases();
  }

  onSelect(release: Release): void {
    this.selectedRelease = release;
  }

  gotoDetail(): void {
    this.router.navigate(['/detail', this.selectedRelease.name]);
  }

  // add(name: string): void {
  //   name = name.trim();
  //   if (!name) { return; }
  //   this.heroService.create(name)
  //     .then(hero => {
  //       this.heroes.push(hero);
  //       this.selectedHero = null;
  //     });
  // }

  delete(name: string): void {
    name = name.trim();
    if (!name) { return; }
    this.releaseService.delete(name)
      .then(response => {
        this.releases = this.releases.filter(rel => rel.name !== name)
      });
  }
}
