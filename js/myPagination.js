function myPagination(_ref) {
    var pageSize = _ref.pageSize,
        pageTotal = _ref.pageTotal,
        curPage = _ref.curPage,
        id = _ref.id,
        getPage = _ref.getPage,
        showPageTotalFlag = _ref.showPageTotalFlag,
        showSkipInputFlag = _ref.showSkipInputFlag,
        pageAmount = _ref.pageAmount,
        dataTotal = _ref.dataTotal,
        blogger = _ref.blogger;

    this.pageSize = pageSize || 5; //分页个数
    this.pageTotal = pageTotal; //总共多少页
    this.pageAmount = pageAmount; //每页多少条
    this.dataTotal = dataTotal; //总共多少数据
    this.curPage = curPage || 1; //初始页码
    this.ul = document.createElement('ul');
    this.id = id;
    this.getPage = getPage;
    this.showPageTotalFlag = showPageTotalFlag || false; //是否显示数据统计
    this.showSkipInputFlag = showSkipInputFlag || false; //是否支持跳转
    this.blogger = blogger;
    console.log("set blogger");
    this.init();
};

// 给实例对象添加公共属性和方法
myPagination.prototype = {
    init: function init() {
        var pagination = document.getElementById(this.id);
        pagination.innerHTML = '';
        this.ul.innerHTML = '';
        pagination.appendChild(this.ul);
        var that = this;
        //首页
        that.firstPage();
        //上一页
        that.lastPage();
        //分页
        that.getPages().forEach(function (item) {
            var li = document.createElement('li');
            if (item == that.curPage) {
                li.className = 'active';
            } else {
                li.onclick = function () {
                    that.curPage = parseInt(this.innerHTML);
                    that.init();
                    that.getPage(that.curPage);
                };
            }
            li.innerHTML = item;
            that.ul.appendChild(li);
        });
        //下一页
        that.nextPage();
        //尾页
        that.finalPage();

        //是否支持跳转
        if (that.showSkipInputFlag) {
            that.showSkipInput();
        }
        //是否显示总页数,每页个数,数据
        if (that.showPageTotalFlag) {
            that.showPageTotal();
        }
    },
    //首页
    firstPage: function firstPage() {
        var that = this;
        var li = document.createElement('li');
        li.innerHTML = '首页';
        if (this.curPage === 1) {
            li.className = 'disabled'; //当前页为首页，禁止跳转至首页
        }
        this.ul.appendChild(li);
        li.onclick = function () {
            if (this.curPage === 1) {
                return;
            }
            var val = parseInt(1);
            that.curPage = val;
            that.getPage(that.curPage);
            that.init();
        };
    },
    //上一页
    lastPage: function lastPage() {
        var that = this;
        var li = document.createElement('li');
        li.innerHTML = '<';
        if (parseInt(that.curPage) > 1) {
            li.onclick = function () {
                that.curPage = parseInt(that.curPage) - 1;
                that.init();
                that.getPage(that.curPage);
            };
        } else {
            li.className = 'disabled';
        }
        this.ul.appendChild(li);
    },
    //分页
    getPages: function getPages() {
        var pag = [];
        if (this.curPage <= this.pageTotal) {
            if (this.curPage <= this.pageSize) {
                //当前页数小于显示条数
                var i = Math.min(this.pageSize, this.pageTotal);
                while (i) {
                    pag.unshift(i--);
                }
            } else {
                //当前页数大于显示条数
                var middle = this.curPage - Math.floor(this.pageSize / 2),
                    //从哪里开始
                    i = this.pageSize;
                if (middle > this.pageTotal - this.pageSize) {
                    middle = this.pageTotal - this.pageSize + 1;
                }
                while (i--) {
                    pag.push(middle++);
                }
            }
        } else {
            console.error('当前页数不能大于总页数');
        }
        if (!this.pageSize) {
            console.error('显示页数不能为空或者0');
        }
        return pag;
    },
    //下一页
    nextPage: function nextPage() {
        var that = this;
        var li = document.createElement('li');
        li.innerHTML = '>';
        if (parseInt(that.curPage) < parseInt(that.pageTotal)) {
            li.onclick = function () {
                that.curPage = parseInt(that.curPage) + 1;
                that.init();
                that.getPage(that.curPage);
            };
        } else {
            li.className = 'disabled';
        }
        this.ul.appendChild(li);
    },
    //尾页
    finalPage: function finalPage() {
        var that = this;
        var li = document.createElement('li');
        li.innerHTML = '尾页';
        if (this.curPage === that.pageTotal) { //当前页为尾页，禁止跳转至尾页
            li.className = 'disabled';
        }
        this.ul.appendChild(li);
        li.onclick = function () {
            var yyfinalPage = that.pageTotal;
            var val = parseInt(yyfinalPage);
            that.curPage = val;
            that.getPage(that.curPage);
            that.init();
        };
    },
    //是否支持跳转
    showSkipInput: function showSkipInput() {
        var that = this;
        var li = document.createElement('li');
        li.className = 'totalPage';
        var span1 = document.createElement('span');
        span1.setAttribute("class", "fl");
        span1.innerHTML = '跳转到';
        li.appendChild(span1);
        var input = document.createElement('input');
        input.setAttribute("type", "number");
        input.onkeydown = function (e) {
            var oEvent = e || event;
            if (oEvent.keyCode == '13') {
                var val = parseInt(input.value);
                if (typeof val === 'number' && val <= that.pageTotal && val > 0) {
                    that.curPage = val;
                    that.getPage(that.curPage);
                } else {
                    swal("跳转页数必须大于等于1，小于等于最大页数!");
                }
                that.init();
            }
        };
        li.appendChild(input);
        var input_bt = document.createElement('input');
        input_bt.setAttribute("type", "button");
        input_bt.setAttribute("id", "page_bt");
        input_bt.value = "确定";
        input_bt.onclick = function () {
            var val = parseInt(input.value);
            if (typeof val === 'number' && val <= that.pageTotal && val > 0) {
                that.curPage = val;
                that.getPage(that.curPage);
            } else {
                swal("跳转页数必须大于等于1，小于等于最大页数!");
            }
            that.init();
        };
        li.appendChild(input_bt);
        // var span2 = document.createElement('span');
        // span2.innerHTML = '页';
        // li.appendChild(span2);
        this.ul.appendChild(li);
    },
    //是否显示总页数,每页个数,数据
    showPageTotal: function showPageTotal() {
        var that = this;
        var li = document.createElement('li');
        li.className = 'totalPage';
        li.innerHTML = '共&nbsp' + that.pageTotal + '&nbsp页&nbsp&nbsp&nbsp' + '每页&nbsp' + that.blogger.pagesize + '&nbsp条&nbsp&nbsp&nbsp'
            + '共&nbsp' + that.blogger.count + '&nbsp条数据';
        this.ul.appendChild(li);
    }
};
function TimestampToDate(Timestamp) {
    let date1 = new Date(Timestamp);
    return date1.toLocaleDateString().replace(/\//g, "-") + " " + date1.toTimeString().substr(0, 8);
}
function GetBlogNode(title, author, address, tags, views, ctime, desc) {
    let node = document.createElement("div");
    node.className = "article-box";
    let abcontent = document.createElement("div");
    abcontent.className = "ab-content";
    let articletitle = document.createElement("div");
    articletitle.className = "article-title";
    let titledoc = document.createElement("a");
    titledoc.href = address;
    titledoc.textContent = title;
    articletitle.appendChild(titledoc);
    abcontent.appendChild(articletitle);
    if (tags != null && tags.length > 0) {
        let cate = document.createElement("div");
        cate.className = "article-cate";
        for (var i = 0; i < tags.length; i++) {
            let ai = document.createElement("a");
            ai.textContent = tags[i];
            ai.href = "tag.html";
            cate.appendChild(ai);
        }
        abcontent.appendChild(cate);
    }
    if (desc != null) {
        let ddoc = document.createElement("div");
        ddoc.className = "article-detail-box c-666";
        ddoc.textContent = desc;
        abcontent.appendChild(ddoc);
    }
    let tailbox = document.createElement("span");
    //view,author account and ctime
    let articledate = document.createElement("span");
    articledate.className = "article-date c-999";
    articledate.textContent = TimestampToDate(ctime);
    tailbox.appendChild(articledate);
    let articleauthor = document.createElement("span");
    articleauthor.className = "article-author one-line-overflow c-999";
    articleauthor.textContent = author;
    tailbox.appendChild(articleauthor);
    abcontent.appendChild(tailbox);
    node.appendChild(abcontent);
    return node;
}
class Blogger {
    pagesize = 10;
    count = 0;
    constructor(account, id, url) {
        this.id = id;
        this.account = account;
        this.url = url;
        console.log("id is ", this.id, this.url);
    }
    async InitPagination() {
        let totalpage = this.count / this.pagesize;
        if (this.count % this.pagesize > 0) {
            totalpage++;
        }
        let pdoc = document.getElementById("pagination");
        pdoc.innerHTML = '';
        let udoc = document.createElement("ul");
        for (let index = 1; index <= totalpage; index++) {
            let lidoc = document.createElement("li");
            lidoc.textContent = index;
            if (index == 1) {
                lidoc.className = "active";
                this.page_id = index;
            }
            lidoc.id = index;
            udoc.appendChild(lidoc);
        }
        const self = this;
        udoc.addEventListener("click", function (e) {
            let pdoc = document.getElementById(self.page_id);
            console.log("click button", e.target.id, pdoc);
            if (e.target.id != self.page_id) {
                pdoc.className = '';
                let pdoc2 = document.getElementById(e.target.id);
                pdoc2.className = "active";
                self.page_id = e.target.id;
                self.BlogListLocal(self.page_id - 1);
            }


        });
        pdoc.appendChild(udoc);
    }
    BlogListLocal(page) {
        if (page * this.pagesize >= this.count) {
            return;
        }
        let last = this.count;
        if ((page + 1) * this.pagesize < last) {
            last = (page + 1) * this.pagesize;
        }
        let htmlElements = document.getElementById(this.id);
        htmlElements.innerHTML = '';
        for (var i = page * this.pagesize; i < last; i++) {
            let node = GetBlogNode(this.rawdata[i].title, this.rawdata[i].author_id, this.rawdata[i].address, this.rawdata[i].tags, this.rawdata[i].view, this.rawdata[i].ctime, this.rawdata[i].desc);
            htmlElements.appendChild(node);
        }
        return;
    }
    BlogList(author_id, tag = "", page = 0, title = "") {
        console.log("do blog list");
        if (this.rawdata != null && this.rawdata.length > 0) {

        }
        fetch(this.url + "/v1/blog/list", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Host": "127.0.0.1:5500"
            },
            body: JSON.stringify({
                "tag": tag, "page": page,
                "page_size": this.pagesize,
                "title": title,
                "author_id": author_id,
            })
        }).then((response) => {
            if (response.status == 200) {
                response.json().then((data) => {
                    if (data.status == 1 && data.data != null) {
                        this.rawdata = data.data;
                        this.count = data.data.length;
                        this.InitPagination();
                        let maxsize = this.pagesize;
                        if (maxsize > data.data.length) {
                            maxsize = data.data.length;
                        }
                        console.log("set count to ", this.count, data.data.length);
                        console.log("data count", data.data.length);
                        let htmlElements = document.getElementById(this.id);
                        for (var i = 0; i < maxsize; i++) {
                            let node = GetBlogNode(data.data[i].title, data.data[i].author_id, data.data[i].address, data.data[i].tags, data.data[i].view, data.data[i].ctime, data.data[i].desc);
                            htmlElements.appendChild(node);
                        }
                        //update page
                    }
                });
            } else {
                console.log("request failed ", response.status, response.body);
            }
        });
        // htmlElements.appendChild(node);
    }

    //get author basic information
    Profile() {
        fetch(this.url + "/v1/user/" + this.account).then((response) => {
            if (response.status == 200) {
                response.json().then((data) => {
                    if (data.status == 1 && data.data != null) {
                        this.uid = data.data.id;
                        this.location = data.data.location;
                        let desc = document.getElementById("person-desc");
                        desc.textContent = data.data.desc;
                        let localations = document.getElementById("person-location");
                        localations.textContent = data.data.location;
                        let account = document.getElementById("account");
                        account.textContent = this.account;
                    }
                });
            } else {
                console.log("request failed ", response.status, response.body);
            }
        });
    }
}

//tagrank
class TagPagination {
    constructor(account, url, author_id, blogid = "article-holder", id = "like-box") {
        this.account = account;
        this.url = url;
        this.author_id = author_id;
        this.doc = document.getElementById(id);
        this.pagesize = 10;
        this.blogdoc = document.getElementById(blogid);
    }
    async InitPagination() {
        let totalpage = this.count / this.pagesize;
        if (this.count % this.pagesize > 0) {
            totalpage++;
        }
        let pdoc = document.getElementById("pagination");
        pdoc.innerHTML = '';
        let udoc = document.createElement("ul");
        for (let index = 1; index <= totalpage; index++) {
            let lidoc = document.createElement("li");
            lidoc.textContent = index;
            if (index == 1) {
                lidoc.className = "active";
                this.page_id = index;
            }
            lidoc.id = index;
            udoc.appendChild(lidoc);
        }
        const self = this;
        udoc.addEventListener("click", function (e) {
            let pdoc = document.getElementById(self.page_id);
            console.log("click button", e.target.id, pdoc);
            if (e.target.id != self.page_id) {
                pdoc.className = '';
                let pdoc2 = document.getElementById(e.target.id);
                pdoc2.className = "active";
                self.page_id = e.target.id;
                self.DeployBlogLocal(self.page_id - 1);
            }


        });
        pdoc.appendChild(udoc);
    }
    getTagNode(name, count) {
        let node = document.createElement("li");
        node.className = "column-category";
        let anode = document.createElement("a");
        anode.innerHTML = `${name}&nbsp;&nbsp;${count}&nbsp;`;
        anode.id = name;
        // let spanode = document.createElement("span");
        node.appendChild(anode);
        return node;
    }
    GetTagList() {
        fetch(this.url + "/v1/blog/tag/list?author_account=" + this.account).then((response) => {
            if (response.status != 200) {
                console.log("bad response", response.body);
                return;
            }
            response.json().then((data) => {
                if (data.status != 1) {
                    console.log("bad response ", data.status);
                    return;
                }
                if (data.data == null || data.data.length == 0) {
                    return;
                }
                this.doc.innerHTML = `<li class="column-title">
                <span class="at-sort b-b-ece fl"><a class="at-sort-comment-a c-666 fl">Tag</a></span>
            </li>`;
                const self = this;
                this.doc.addEventListener("click", function (e) {
                    console.log("click", e.target);
                    self.GetBlogWithTag(e.target.id);
                });
                for (let index = 0; index < data.data.length; index++) {
                    this.doc.appendChild(this.getTagNode(data.data[index].key, data.data[index].value));
                }
            });
        });
    }
    DeployBlogLocal(page) {
        let last = this.rawdata.length;
        if (last - page * this.pagesize > this.pagesize) {
            last = (page + 1) * this.pagesize;
        }
        this.blogdoc.innerHTML = '';
        for (let i = page * this.pagesize; i < last; i++) {
            let bnode = GetBlogNode(this.rawdata[i].title, this.rawdata[i].author_id, this.rawdata[i].address, this.rawdata[i].tags, this.rawdata[i].view, this.rawdata[i].ctime, this.rawdata[i].desc);
            this.blogdoc.appendChild(bnode);
        }
    }
    GetBlogWithTag(tagname) {
        fetch(this.url + "/v1/blog/list", {
            method: "POST", body: JSON.stringify({
                "author_id": this.author_id,
                "tag": tagname
            })
        }).then((response) => {
            if (response.status != 200) {
                console.log("bad response", response.body);
                return;
            }
            response.json().then((data) => {
                if (data.status != 1) {
                    console.log("bad response ", data.status);
                    return;
                }
                if (data.data == null || data.data.length == 0) {
                    return;
                }
                this.count = data.data.length;
                this.InitPagination();
                this.rawdata = data.data;
                this.DeployBlogLocal(0);
            });
        });
    }
}

class User {
    Login() {
    }
    IsLogin() {

        return false;
    }
}