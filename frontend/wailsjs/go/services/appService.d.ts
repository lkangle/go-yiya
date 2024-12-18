// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';
import {context} from '../models';

export function GetAppInfo():Promise<types.JsObject>;

export function GetCompressOptions():Promise<types.JSResp>;

export function GetIsAlwaysOnTop():Promise<boolean>;

export function GetLocalPosition():Promise<number|number>;

export function GetWindowSize():Promise<number|number>;

export function SaveBoundary():Promise<void>;

export function SetIsAlwaysOnTop(arg1:boolean):Promise<void>;

export function Setup(arg1:context.Context):Promise<void>;

export function UpdateCompressOptions(arg1:types.CompressOptions):Promise<types.JSResp>;
