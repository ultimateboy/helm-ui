import { Component, OnInit, Input } from '@angular/core';
import { Router }            from '@angular/router';

import { Observable }        from 'rxjs/Observable';
import { Subject }           from 'rxjs/Subject';

// Observable class extensions
import 'rxjs/add/observable/of';

import { Release } from './release';
import { ReleaseService } from './release.service';


@Component({
  selector: 'release-notes',
  templateUrl: './release-notes.component.html',
  styleUrls: [ './release-notes.component.css' ]
})
export class ReleaseNotesComponent implements OnInit {
  @Input() releaseName: string;
  release: Release;
  show: boolean;

  constructor(
    private releaseService: ReleaseService
  ) { }

  toggleShow(): void {
    this.show = this.show ? false:true;
  }

  ngOnInit(): void {
    this.getRelease(this.releaseName);
  }

  getRelease(name: string): void {
    this.releaseService.getRelease(name).then(release => this.release = release);
  }

}
