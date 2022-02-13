now=$(date +%s)
ffmpeg -i /video/hls/$1.m3u8 -c copy -bsf:a aac_adtstoasc /video/vod/$1/$now.mp4
