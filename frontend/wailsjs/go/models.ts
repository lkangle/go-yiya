export namespace types {
	
	export class CompressOptions {
	    quality: number;
	    override: boolean;
	    newSuffix: string;
	
	    static createFrom(source: any = {}) {
	        return new CompressOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quality = source["quality"];
	        this.override = source["override"];
	        this.newSuffix = source["newSuffix"];
	    }
	}
	export class CompressResult {
	    $input_path: string;
	    $input_temp_path: string;
	    path: string;
	    size: number;
	    code: number;
	    success: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new CompressResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.$input_path = source["$input_path"];
	        this.$input_temp_path = source["$input_temp_path"];
	        this.path = source["path"];
	        this.size = source["size"];
	        this.code = source["code"];
	        this.success = source["success"];
	        this.message = source["message"];
	    }
	}
	export class ImageFileInfo {
	    id: string;
	    dir_path: string;
	    filename: string;
	    path: string;
	    size: number;
	    contentType: string;
	
	    static createFrom(source: any = {}) {
	        return new ImageFileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.dir_path = source["dir_path"];
	        this.filename = source["filename"];
	        this.path = source["path"];
	        this.size = source["size"];
	        this.contentType = source["contentType"];
	    }
	}
	export class JSResp {
	    success: boolean;
	    code: number;
	    message: string;
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new JSResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}

}

