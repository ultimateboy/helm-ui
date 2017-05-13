import { Component, OnInit, Input, Optional, Inject,Output,EventEmitter } from '@angular/core';
import { Router }            from '@angular/router';
import { MdDialog, MdDialogRef, MD_DIALOG_DATA } from '@angular/material';

import { Observable }        from 'rxjs/Observable';
import { Subject }           from 'rxjs/Subject';

// Observable class extensions
import 'rxjs/add/observable/of';

import { Release } from './release';
import { ReleaseService } from './release.service';


@Component({
  selector: 'release-controls',
  templateUrl: './release-controls.component.html',
  styleUrls: [ './release-controls.component.css' ]
})
export class ReleaseControlsComponent implements OnInit {
  @Input() releaseName: string;
  @Input() ParentReleases: Release[];
  @Output() outputEvent:EventEmitter<string>=new EventEmitter();
  oldReleases: Release[];
  release: Release;
  showNotes: boolean;
  showHistory: boolean;
  dialogResp: string;
  loading: boolean;

  constructor(
    private releaseService: ReleaseService,
    private _dialog: MdDialog
  ) { }

  ngOnInit(): void {
    this.getRelease(this.releaseName);
    this.getReleaseHistory(this.releaseName);
  }

  getRelease(name: string): void {
    this.releaseService.getRelease(name).then(release => this.release = release);
  }
  getReleaseHistory(name: string): void {
    this.releaseService.getReleaseHistory(name)
      .then(releases => {
        this.oldReleases = releases.reverse();
      });
  }
  delete(name: string): void {
    name = name.trim();
    if (!name) { return; }
    this.releaseService.delete(name)
      .then(response => {
        this.outputEvent.emit(name);
      });
  }
  rollback(name: string, revision: number): void {
    if (!name || !revision) { return; }
    this.releaseService.rollback(name, revision)
      .then(release => {
        console.log(release);
      });
  }
  openEditDialog(rel: Release) {
    let configData = rel.config.raw ? rel.config.raw.trim():"";
    const dialogRef = this._dialog.open(DialogContentComponent, {
      data: {'config':configData, 'values':rel.chart.values.raw},
    });
    dialogRef.afterClosed().subscribe(result => {
      this.dialogResp = result;
      if (result) {
        this.toggleLoad();

        this.releaseService.updateValues(rel.name, result)
          .then(release => {
            for (var i = 0; i < this.ParentReleases.length; i++) {
              if (this.ParentReleases[i].name == release.name) {
                this.ParentReleases[i] = release;
              }
            }
          });
      }
    })
  }

  toggleLoad(): void {
    this.loading = this.loading ? false: true; 
  }
  toggleShowHistory(): void {
    this.showHistory = this.showHistory ? false:true;
  }
  toggleShowNotes(): void {
    this.showNotes = this.showNotes ? false:true;
  }
}



@Component({
  template: `
    <textarea #editorInput>
{{ data.config }}
    </textarea>
    <pre>
{{ data.values }}
    </pre>
    <br />
    <button color="accent" md-button (click)="dialogRef.close(editorInput.value)">
      <md-icon>save</md-icon> save
    </button>
    <button color="accent" md-button (click)="dialogRef.close()">
      <md-icon>cancel</md-icon> cancel
    </button>
  `,
  styles: [`
    textarea {
      width: 50em;
      height: 10em;
    }
    pre {
      width: 50em;
      height: 10em;
      overflow: auto;
      background-color: #eeeeee;
      word-break: normal !important;
      word-wrap: normal !important;
      white-space: pre !important;
    }
  `],
})
export class DialogContentComponent {
  code: string;
  constructor( 
    @Optional() public dialogRef: MdDialogRef<DialogContentComponent>,
    @Inject(MD_DIALOG_DATA) public data: any
  ) { }
}

import { Pipe, PipeTransform } from '@angular/core';

@Pipe({name: 'unepoch'})
export class UnEpochPipe implements PipeTransform {
  transform(value: number, args: string[]): any {
    if (!value) return value;
    var t = new Date(1970, 0, 1); // Epoch
    t.setSeconds(value);

    return t.toLocaleDateString("en-US");
  }
}

@Pipe({name: 'status'})
export class StatusStringPipe implements PipeTransform {
  transform(value: number, args: string[]): any {
    let statuses = [
      "UNKNOWN",
      "DEPLOYED",
      "DELETED",
      "SUPERSEDED",
      "FAILED",
      "DELETING"
    ];
    
    if (!value) return value;
    if (value > statuses.length) return "UNKNOWN";
    
    return statuses[value];
  }
}
