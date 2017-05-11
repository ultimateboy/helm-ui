import { Component, OnInit, Optional, Inject } from '@angular/core';
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
    private releaseService: ReleaseService
  ) { }

  getReleases(): void {
    this.releaseService.getReleases().then(releases => this.releases = releases);
  }

  ngOnInit(): void {
    this.getReleases();
  }


  onSelect(release: Release): void {
    this.selectedRelease = release;
  }
  onComponentChange(value: string){
   console.log("I have a values!!!" + value);
   this.releases = this.releases.filter(rel => rel.name !== value)

  }

}

