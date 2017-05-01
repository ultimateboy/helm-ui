import { Injectable }    from '@angular/core';
import { Headers, Http } from '@angular/http';
import { APIURL } from './config/config';

import 'rxjs/add/operator/toPromise';

import { ChartRepo } from './chart-repo';
import { Chart } from './chart';
import { Release } from './release';

@Injectable()
export class ChartRepoService {

  private reposUrl = APIURL + '/repos';

  constructor(private http: Http) { }

  getRepos(): Promise<ChartRepo[]> {
    return this.http.get(this.reposUrl)
               .toPromise()
               .then(response => response.json() as ChartRepo[])
               .catch(this.handleError);
  }

  getRepoCharts(name: string): Promise<Chart[]> {
    return this.http.get(this.reposUrl+'/'+name+'/charts')
               .toPromise()
               .then(response => response.json() as Chart[])
               .catch(this.handleError);
  }


  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getReposSlowly(): Promise<ChartRepo[]> {
    return new Promise(resolve => {
      // Simulate server latency with 2 second delay
      setTimeout(() => resolve(this.getRepos()), 2000);
    });
  }

  private headers = new Headers({'Content-Type': 'application/json'});

  create(name: string, url: string): Promise<ChartRepo> {
    return this.http
      .post(this.reposUrl, JSON.stringify({name: name, url: url}), {headers: this.headers})
      .toPromise()
      .then(res => res.json() as ChartRepo)
      .catch(this.handleError);
  }

  install(chart: string, repo: string): Promise<Release> {
    return this.http
      .post(this.reposUrl+'/'+repo+'/charts/'+chart+'/install', JSON.stringify({name: name}), {headers: this.headers})
      .toPromise()
      .then(res => res.json() as Release)
      .catch(this.handleError);
  }

  search(repo: string, term: string): Promise<Chart[]> {
    return this.http.get(this.reposUrl+'/'+repo+'/charts?name='+ term)
               .toPromise()
               .then(response => response.json() as Chart[])
               .catch(this.handleError);
  }


}
