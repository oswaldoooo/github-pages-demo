const VIEW = 1;//文章实时阅读数据
const COMMENT = 2;//文章实时评论
// 文章实时互动
class Article {
  constructor(baseurl, articleid) {
    this.url = baseurl + "?" + articleid;
    this.conn = new WebSocket(this.url);
    this.conn.onerror = function (event) {
      console.log("remote closed ", event);
    }
    this.view = function (data) {
      console.log("accept data", data);
    }
    this.comment = this.view;
    this.conn.onclose = function (event) {
      console.log("connection closed");
    }
    const self = this;
    this.conn.onmessage = function (msg) {
      let data = JSON.parse(msg.data);
      console.log("raw:", data);
      switch (data.code) {
        case VIEW:
          //todo:更新数据
          self.view(data);
          break;
        case COMMENT:
          //本地显示新添加评论
          self.comment(data);
          break;
        default:
          console.log("unknown code ", data.code);
      }
    }
  }
  onview(event) {
    this.view = event;
  }
  oncomment(event) {
    this.comment = event;
  }
}