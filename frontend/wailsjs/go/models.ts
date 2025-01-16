export namespace types {
	
	export class CompletionConfigRequest {
	    api: string;
	    sk: string;
	    model: string;
	    adapter: string;
	    maxTokens: number;
	
	    static createFrom(source: any = {}) {
	        return new CompletionConfigRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.api = source["api"];
	        this.sk = source["sk"];
	        this.model = source["model"];
	        this.adapter = source["adapter"];
	        this.maxTokens = source["maxTokens"];
	    }
	}
	export class GuiResponse {
	    code: number;
	    data: any;
	    msg: string;
	
	    static createFrom(source: any = {}) {
	        return new GuiResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.data = source["data"];
	        this.msg = source["msg"];
	    }
	}
	export class SystemConfigRequest {
	    domain: string;
	    port: number;
	    sslCert: string;
	    sslKey: string;
	
	    static createFrom(source: any = {}) {
	        return new SystemConfigRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.domain = source["domain"];
	        this.port = source["port"];
	        this.sslCert = source["sslCert"];
	        this.sslKey = source["sslKey"];
	    }
	}

}

