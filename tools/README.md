Merge body \<style amp-custom> to head by:

```
find ./ -type f -name "*.html" -exec ./fix_style_for_amp {} \;
```
