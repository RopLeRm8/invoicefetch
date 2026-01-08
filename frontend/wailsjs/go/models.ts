export namespace main {
	
	export class AuthResult {
	    ok: boolean;
	    url: string;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new AuthResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.url = source["url"];
	        this.error = source["error"];
	    }
	}

}

