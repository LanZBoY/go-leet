#!/usr/bin/env python3
"""抓 LeetCode 題目原文，輸出 Markdown 到 problems/<dir>/README.md

用法:
    python3 scripts/fetch-problem.py p0125_valid_palindrome
    python3 scripts/fetch-problem.py p0125_valid_palindrome valid-palindrome  # 指定 slug
    python3 scripts/fetch-problem.py --all                                    # 回填所有題目

slug 預設由資料夾名推導：p0125_valid_palindrome -> valid-palindrome
"""
import html
import json
import os
import re
import sys
import urllib.request

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
API = "https://leetcode.com/graphql/"
QUERY = """query q($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionFrontendId title difficulty content
    topicTags { name }
  }
}"""


def fetch(slug):
    req = urllib.request.Request(
        API,
        data=json.dumps({"query": QUERY, "variables": {"titleSlug": slug}}).encode(),
        headers={
            "Content-Type": "application/json",
            "User-Agent": "Mozilla/5.0",
            "Referer": f"https://leetcode.com/problems/{slug}/",
        },
    )
    with urllib.request.urlopen(req, timeout=20) as r:
        return json.load(r)["data"]["question"]


def to_markdown(h):
    """LeetCode 的 content 是 HTML，轉成堪讀的 Markdown。"""
    h = re.sub(r"<sup>(.*?)</sup>", r"^\1", h, flags=re.S)
    h = re.sub(r"<sub>(.*?)</sub>", r"_\1", h, flags=re.S)

    # <pre> 內是範例輸入輸出，整塊當 code fence，先抽出來保護
    blocks = []

    def stash(m):
        inner = re.sub(r"<[^>]+>", "", m.group(1))
        blocks.append(html.unescape(inner).strip("\n"))
        return f"\x00{len(blocks) - 1}\x00"

    h = re.sub(r"<pre>(.*?)</pre>", stash, h, flags=re.S)

    h = re.sub(r"<code>(.*?)</code>", r"`\1`", h, flags=re.S)
    h = re.sub(r"<(strong|b)[^>]*>(.*?)</\1>", r"**\2**", h, flags=re.S)
    # LeetCode 常用 <em> 把半句話包起來（如 return `true`<em> if ...</em>），
    # 直譯成 * 會變成破碎的星號，所以直接去掉標記留文字。
    h = re.sub(r"</?(em|i)>", "", h)
    h = re.sub(r"<li>(.*?)</li>", lambda m: "- " + " ".join(m.group(1).split()) + "\n", h, flags=re.S)
    h = re.sub(r"</?(ul|ol)>", "\n", h)
    h = re.sub(r"<br\s*/?>", "\n", h)
    h = re.sub(r"</p>", "\n\n", h)
    h = re.sub(r"<img[^>]*src=\"([^\"]+)\"[^>]*>", r"![](\1)", h)
    h = re.sub(r"<[^>]+>", "", h)
    h = html.unescape(h).replace("\xa0", " ")

    # 去掉 HTML 原始碼縮排造成的 tab（只動清單行，避免破壞 code block）
    h = re.sub(r"^[ \t]+- ", "- ", h, flags=re.M)
    h = re.sub(r"(?m)(^- .*\n)\n+(?=- )", r"\1", h)

    for i, b in enumerate(blocks):
        h = h.replace(f"\x00{i}\x00", f"```\n{b}\n```")

    h = re.sub(r"[ \t]+\n", "\n", h)
    h = re.sub(r"\n{3,}", "\n\n", h)
    return h.strip()


def write(dirname, slug=None):
    slug = slug or dirname.split("_", 1)[1].replace("_", "-")
    q = fetch(slug)
    if not q:
        print(f"  !! {dirname}: slug '{slug}' 查無此題")
        return False
    tags = ", ".join(t["name"] for t in q["topicTags"])
    md = (
        f"# {q['questionFrontendId']}. {q['title']}\n\n"
        f"- **Difficulty**: {q['difficulty']}\n"
        f"- **Topics**: {tags}\n"
        f"- **Link**: https://leetcode.com/problems/{slug}/\n\n"
        f"---\n\n{to_markdown(q['content'])}\n"
    )
    path = os.path.join(ROOT, "problems", dirname, "README.md")
    with open(path, "w") as f:
        f.write(md)
    print(f"  ok {dirname} ({len(md)} bytes)")
    return True


if __name__ == "__main__":
    if len(sys.argv) > 1 and sys.argv[1] == "--all":
        base = os.path.join(ROOT, "problems")
        for d in sorted(os.listdir(base)):
            if d.startswith("p") and os.path.isdir(os.path.join(base, d)):
                try:
                    write(d)
                except Exception as e:
                    print(f"  !! {d}: {e}")
    elif len(sys.argv) > 1:
        write(sys.argv[1], sys.argv[2] if len(sys.argv) > 2 else None)
    else:
        print(__doc__)
        sys.exit(1)
