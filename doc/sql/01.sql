
SELECT *
from
    qy_blog_comment_agent

-- 更新评论agnet数量
update    qy_blog_comment_agent a set `count` =  (
    select count(*)
    from qy_blog_comment_index qbci
    where qbci.deleted_flag  = '0'
      and qbci.agent_id = a.id
),root_count =  (
    select count(*)
    from qy_blog_comment_index qbci
    where qbci.deleted_flag  = '0'
      and qbci.parent_id =0
      and qbci.agent_id = a.id
)
-- 更新子评论总数
update  qy_blog_comment_index a set `count` =  (
    select num
    from
        (
            select count(*) num,a2.root_id
            from qy_blog_comment_index a2
            where
                    a2.deleted_flag  = '0'
            GROUP BY a2.root_id
        ) c
    where  root_id =   a.id
)
update  qy_blog_comment_index a set `count` =  0 where `count`  is NULL

update  qy_blog_comment_index a set root_count  =  (
    select num
    from
        (
            select count(*) num,a2.root_id
            from qy_blog_comment_index a2
            where
                    a2.deleted_flag  = '0'
              AND a2.parent_id = a2.root_id
            GROUP BY a2.root_id
        ) c
    where  root_id =   a.id
)

update  qy_blog_comment_index a set root_count =  0 where root_count  is NULL

-- 更新文章评论总数
UPDATE qy_blog_article  a set comment_count = (
    select `count`
    from qy_blog_comment_agent ca
    where ca.id  = a.comment_agent_id
)



update    qy_blog_comment_agent a set root_count =  (
    select count(*)
    from qy_blog_comment_index qbci
    where qbci.deleted_flag  = '0'
      and qbci.parent_id =0
      and qbci.agent_id = a.id
)

