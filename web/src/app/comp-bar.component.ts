import { Component, OnInit, Input } from '@angular/core';
import { Router }            from '@angular/router';

import { Observable }        from 'rxjs/Observable';
import { Subject }           from 'rxjs/Subject';

// Observable class extensions
import 'rxjs/add/observable/of';

import { CTemplate } from './release';


export class CBlock {
  resource: string;
  size: number;
}


@Component({
  selector: 'comp-bar',
  templateUrl: './comp-bar.component.html',
  styleUrls: [ './comp-bar.component.css' ]
})
export class CompBarComponent implements OnInit {
  @Input() templates: CTemplate[];
  name: string; 
  blocks: CBlock[];
  resources: string[];

  constructor() { 
    this.blocks = [];
    this.resources = [];
  }

  ngOnInit(): void {
    for (var i = 0; i < this.templates.length; i++) {
      var decoded = atob(this.templates[i].data);
      var matches = decoded.match(/kind\: (\w+)/);
      if (matches) {
        this.resources.push(matches[1]);
      }
    }

    this.resources.sort();

    var counts = {};
    for(var i = 0; i< this.resources.length; i++) {
        var k = this.resources[i];
        counts[k] = counts[k] ? counts[k]+1 : 1;
    }

    for (k in counts) {
      var blockCount = (counts[k]/this.resources.length) * 100;
      var b = new CBlock;
      b.resource = k;
      b.size = blockCount;
      this.blocks.push(b)
    }
    
  }
}
