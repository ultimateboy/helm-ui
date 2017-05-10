import { Injectable }    from '@angular/core';
import { Headers, Http } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import { APIURL } from './config/config';

import { Response } from './response';
import { Release } from './release';

@Injectable()
export class ReleaseService {

  private releasesUrl = APIURL + '/releases';

  constructor(private http: Http) { }

  getReleases(): Promise<Release[]> {
    return this.http.get(this.releasesUrl)
               .toPromise()
               .then(response => response.json() as Release[])
               .catch(this.handleError);
  }

  getRelease(name: string): Promise<Release> {
    return this.http.get(this.releasesUrl + "/" + name)
               .toPromise()
               .then(response => response.json() as Release)
               .catch(this.handleError);
  }

  getChartReleases(name: string): Promise<Release[]> {
    return this.http.get(this.releasesUrl+"?chart="+name)
               .toPromise()
               .then(response => response.json() as Release[])
               .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getReposSlowly(): Promise<Release[]> {
    return new Promise(resolve => {
      // Simulate server latency with 2 second delay
      setTimeout(() => resolve(this.getReleases()), 2000);
    });
  }

  private headers = new Headers({'Content-Type': 'application/json'});

  delete(name: string): Promise<Response> {
    return this.http
      .delete(this.releasesUrl + '/'+name)
      .toPromise()
      .then(res => res.json() as Response)
      .catch(this.handleError);
  }

  updateValues(name: string, data: string): Promise<Release> {
    return this.http
      .patch(this.releasesUrl + '/'+name, JSON.stringify({data: data}), {headers: this.headers})
      .toPromise()
      .then(res => res.json() as Release)
      .catch(this.handleError);
  }

}
