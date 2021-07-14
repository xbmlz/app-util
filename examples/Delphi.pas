unit DelphiEx;

(**
 * @desc   delphi 扩展单元
 * @author chenxc
 * @date   2020年11月30日17:07:43
 * @note   注意：需要在工程文件(.dpr)和引用此单元的文件(.pas)中的uses最前面添加ShareMem
 *)

interface

(**
 * @desc   get请求
 * @param  url      请求地址
 * @param  resp     http返回值(指针类型)
 * @param  headers  请求头，默认为nil. eg: {"a":"abc", "b": "cba"}
 * @param  proxyUrl 代理路径
 * @return 0 成功，-1失败
 *)
function HttpGet(url: PChar; var resp: PChar; headers: PChar = nil; proxyUrl: PChar = nil): Integer; cdecl; external 'DelphiEx.dll';

(**
 * @desc   post请求(content-type: form-data)
 * @param  url      请求地址
 * @param  data     请求参数。 eg: {"a":"abc", "b": "cba"}
 * @param  resp     http返回值(指针类型)
 * @param  headers  请求头，默认为nil. eg: {"a":"abc", "b": "cba"}
 * @param  proxyUrl 代理路径
 * @return 0 成功，-1失败
 *)
function HttpPost(url: PChar; data: PChar; var resp: PChar; headers: PChar = nil; proxyUrl: PChar = nil): Integer; cdecl; external 'DelphiEx.dll';

(**
 * @desc   post请求(content-type: application/json)
 * @param  url      请求地址
 * @param  data     请求参数。 eg: {"a":"abc", "b": "cba"}
 * @param  resp     http返回值(指针类型)
 * @param  headers  请求头，默认为nil. eg: {"a":"abc", "b": "cba"}
 * @param  proxyUrl 代理路径
 * @return 返回值      0 成功，-1: headers错误， -2: 请求参数格式化失败， -3: Http请求失败
 *)
function HttpPostJson(url: PChar; data: PChar; var resp: PChar; headers: PChar = nil; proxyUrl: PChar = nil): Integer; cdecl; external 'DelphiEx.dll';

(**
 * @desc   上传文件
 * @param  url          请求地址
 * @param  fieldName    字段名称 eg：file
 * @param  filePath     文件地址 eg：E://a.jpg
 * @param  resp         http返回值(指针类型)
 * @param  data         请求参数，默认为nil. eg: {"a":"abc", "b": "cba"}
 * @param  headers      请求头，默认为nil. eg: {"a":"abc", "b": "cba"}
 * @param  proxyUrl     代理路径
 * @return 返回值       0 成功，-1: headers错误，-2：参数错误， -3: Http请求失败
 *)
function UploadFile(url: PChar; fieldName: PChar; filePath: PChar; var resp: PChar; data: PChar = nil; headers: PChar = nil; proxyUrl: PChar = nil): Integer; cdecl; external 'DelphiEx.dll';

(**
 * @desc   下载文件
 * @param  url          请求地址
 * @param  filename     文件保存名字 eg：E://test.jpg
 * @param  headers      请求头，默认为nil. eg: {"a":"abc", "b": "cba"}
 * @param  proxyUrl     代理路径
 * @return 返回值       0 成功，-1: headers错误， -3: Http请求失败
 *)
function DownloadFile(url: PChar; filename: PChar; headers: PChar = nil; proxyUrl: PChar = nil): Integer; cdecl; external 'DelphiEx.dll';

implementation

end.

