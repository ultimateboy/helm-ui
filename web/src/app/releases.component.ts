import { Component, OnInit, Optional, Inject } from '@angular/core';
import { Router } from '@angular/router';
import { MdDialog, MdDialogRef, MD_DIALOG_DATA } from '@angular/material';

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
    private releaseService: ReleaseService,
    private _dialog: MdDialog
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

  openDialog(config: string) {
    const dialogRef = this._dialog.open(DialogContentComponent, {
      data: config,
    });

  }

}


@Component({
  template: `
    <pre>
      {{ data }}
    </pre>
    <button md-button (click)="dialogRef.close()">CLOSE</button>
  `,
})
export class DialogContentComponent {
  constructor( 
    @Optional() public dialogRef: MdDialogRef<DialogContentComponent>,
    @Inject(MD_DIALOG_DATA) public data: any
  ) { }
}