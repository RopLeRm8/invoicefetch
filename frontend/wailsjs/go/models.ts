export namespace main {
	
	export class ActiveEmail {
	    email: string;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new ActiveEmail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.email = source["email"];
	        this.error = source["error"];
	    }
	}
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
	export class Logout {
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new Logout(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	    }
	}
	export class Verify {
	    email?: string;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new Verify(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.email = source["email"];
	        this.error = source["error"];
	    }
	}

}

